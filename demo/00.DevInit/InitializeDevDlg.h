#ifndef INITIALIZEDEVDLG_H
#define INITIALIZEDEVDLG_H

#include <QDialog>

namespace Ui {
class CInitializeDevDlg;
}

class CInitializeDevDlg : public QDialog
{
    Q_OBJECT
    
public:
    explicit CInitializeDevDlg(QWidget *parent = 0);
    ~CInitializeDevDlg();

    //初始化函数
    void Init();
    //设置重置方式
    void SetRetWay(const QString& strRetWay);
    //获取密码
    QString GetPwd();
    //获取用户名
    QString GetUser();
    //获取重置方式
    QString GetPwdRestWay();

private slots:
    void on_CancelButton_clicked();

    void on_OKButton_clicked();

protected:
    virtual void keyPressEvent(QKeyEvent *);

private:
    Ui::CInitializeDevDlg *ui;

public:
    QString	m_strRig;
    QString	m_strUserName;
    QString	m_strPwdRestWay;
    QString m_strConfirmPwd;
    QString m_strPwd;
};

#endif // INITIALIZEDEVDLG_H
