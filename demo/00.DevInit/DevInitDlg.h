#ifndef DEVINITDLG_H
#define DEVINITDLG_H

#include <QDialog>
#include <QStandardItemModel>
#include "dhnetsdk.h"

#define MAX_DEV_INFO_COUNT      (1024*32)
#ifndef NEW
#define NEW new(std::nothrow)
#endif

namespace Ui {
class CDevInitDlg;
}

typedef enum
{
    LISTCOLUMN_STATUS = 0,
	LISTCOLUMN_IPVERSION,
    LISTCOLUMN_IPADDRESS,
	LISTCOLUMN_PORT,
	LISTCOLUMN_SUBNETMASK,
	LISTCOLUMN_GATEWAY,
	LISTCOLUMN_MACADDRESS,
	LISTCOLUMN_DEVTYPE,
	LISTCOLUMN_DETAILTYPE,
	LISTCOLUMN_HTTP,
	LISTCOLUMN_COUNT = 10
}LISTVIEW_COLUMN;

class CDevInitDlg : public QDialog
{
    Q_OBJECT
    
public:
    explicit CDevInitDlg(QWidget *parent = 0);
    ~CDevInitDlg();

    friend void CALLBACK cbSearchDevices(DEVICE_NET_INFO_EX *pDevNetInfo, void* pUserData);
    //初始化
    void Init();
    //初始化设备搜索List列表
    void InitListView();
    //初始化SDK
    void InitNetSDK();
    //设备初始化状态
    BOOL GetInitStatus(BYTE initStatus);
    //指定IP搜索
    void StartSearchDeviceByIP(const QString& strStartIP, const QString& strEndIP);
    //IP格式检查
    BOOL CheckIP(const QString& strStartIP, const QString& strEndIP);
    //IP转换
    void IPtoStr(DWORD ip, char* buf, unsigned int nBufferSize);
    //获取密码重置方式
    void GetPwdRestWay(BYTE pwdRestWay);

signals:
    void SearchDevices(DEVICE_NET_INFO_EX* pData);
    void SearchDevicesByIP(const QString& strStartIP, const QString& strEndIP);

private slots:
    void on_InitializeDevice_Button_clicked();

    void on_ByIPSearchButton_clicked();

    void on_Broadcast_Button_clicked();

    void OnSearchDevices(DEVICE_NET_INFO_EX* pDevNetInfo);

    void OnSearchDevicesByIP(const QString& strStartIP, const QString& strEndIP);

protected:
    virtual void keyPressEvent(QKeyEvent *);

private:
    Ui::CDevInitDlg *ui;

private:
    QStandardItemModel *m_Model;
    std::vector<DEVICE_NET_INFO_EX*> m_DevNetInfo;
    int  m_nDeviceCount;           //当前vector中的元素的个数
    LLONG m_lpSearch;              //是否在搜索
    DWORD m_dwStartIP;             //起始IP
    DWORD m_dwEndIP;               //结束IP
    int  m_nSelected;              //当前选中的第几个元素
    QString m_strPwdResetWay;      //密码重置方式

};

#endif // DEVINITDLG_H
