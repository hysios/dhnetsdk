package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"

	dhnetsdk "github.com/hysios/dhnetsdk"
)

func main() {
	client, err := dhnetsdk.ClientInit(func(client *dhnetsdk.Client, ip string, port int) {
		log.Printf("disconnect %v ip %s port %d", client, ip, port)
	})

	if err != nil {
		log.Fatalf("init client error %s", err)
	}

	search := dhnetsdk.NetSearchDevice{}
	ifaces, err := net.Interfaces()
	// handle err
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return
		}
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
				log.Printf("ipv4: %s\n", ip)
			case *net.IPAddr:
				ip = v.IP
				log.Printf("ipv6: %s\n", ip)
			}

			search.SetLocalIP(ip.String())
			search.Start(func(search *dhnetsdk.NetSearchDevice, deviceinfo *dhnetsdk.DeviceNetInfoEx) {
				log.Printf("deviceinfo\n%s\n", deviceinfo.DeviceNetInfo)
			}, dhnetsdk.StMulticastAndBroadcast)
			// process IP address

			time.Sleep(5 * time.Second)
			search.Stop()
		}
	}

	var security dhnetsdk.LoginSecurity
	security.SetIP("192.168.1.108")
	security.SetPort(37777)
	security.SetUserName("admin")
	security.SetPassword("admin123")
	security.SetSpecCap(dhnetsdk.EmLoginSpecCapTcp)

	client.SetAutoReconnect(func(client *dhnetsdk.Client, ip string, port int) {
		log.Printf("重连设备 ip %s port %d", ip, port)
	})

	client.SetDVRMessCallBack(func(client *dhnetsdk.Client, cmd dhnetsdk.DhAlarmType, buf []byte, ip string, port int) bool {
		log.Printf("cmd %d", cmd)
		return true
	})

	log.Printf("ip %s, port %d\n", security.IP(), security.Port())
	log.Printf("admin %s cap %v\n", security.UserName(), security.SpecCap())
	// security.Print()
	if err = client.Login("192.168.1.108:37777", "admin", "admin123"); err != nil {
		log.Fatalln(err)
	}

	log.Printf("serialNumber %s", string(client.DeviceInfo.SerialNumber()))
	log.Printf("Client DeviceInfo\n%s\n", client.DeviceInfo)
	client.RealLoadPictureEx(0, dhnetsdk.EventIvsAll, func(client *dhnetsdk.Client, alarmType dhnetsdk.EventIvs, alarmInfo interface{}, frame []byte, seq int) int {
		switch alarmType {
		case dhnetsdk.EventIvsTrafficParking,
			dhnetsdk.EventIvsTrafficParkingB,
			dhnetsdk.EventIvsTrafficParkingC,
			dhnetsdk.EventIvsTrafficParkingD:
			if info, ok := alarmInfo.(*dhnetsdk.TrafficParkingInfo); ok {
				log.Printf("TrafficParkingInfo\n%s\n", info)
				log.Printf("TrafficParkingInfo.Object\n%s\n", info.Object())
				log.Printf("TrafficParkingInfo.Vehicle\n%s\n", info.Vehicle())
				vehi := info.Vehicle()
				log.Printf("Vehicle.Color\n%s\n", vehi.MainColor())

				obj := info.Object()
				log.Printf("Object.BoundingBox\n%s\n", obj.BoundingBox())
				log.Printf("Object.Center\n%s\n", obj.Center())
				log.Printf("Object.OriginalBoundingBox\n%s\n", obj.OriginalBoundingBox())

			}

			t := time.Now()
			log.Printf("frame size %d seq %d\n", len(frame), seq)
			if err := ioutil.WriteFile(fmt.Sprintf("tmp/%s-%d.jpg", t.Format("20060102150405.999"), seq), frame, os.ModePerm); err != nil {
				log.Printf("write error %s", err)
			}
		}
		return 0
	})

	// var mediaQuery = dhnetsdk.MediaQueryTrafficcarEx{}
	// mediaQuery.Init()
	// mediaQuery.Param.SetMediaType(1)
	// mediaQuery.Param.SetChannelID(0)
	// mediaQuery.Param.SetStartTime(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local))
	// mediaQuery.Param.SetEndTime(time.Now())
	// // mediaQuery.SetFileFlag(0)
	// // mediaQuery.Param.SetFileFlagEx(0x12)

	// mediaQuery.Param.SetEventTypes([]dhnetsdk.EventIvs{dhnetsdk.EventIvsTrafficParking})

	for {
		time.Sleep(1 * time.Second)
	}
	// findfile.Close()

	dhnetsdk.ClientCleanup()

}
