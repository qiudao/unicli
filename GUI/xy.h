#ifndef XY_H
#define XY_H

#include <QWidget>
#include <QLabel>
#include <QPushButton>
#include <QLineEdit>

class XY : public QWidget
{
	Q_OBJECT

public:
	XY(QWidget *parent = 0);
	~XY();

	float ymin() {return m_ymin;}
	float ymax() {return m_ymax;}
	int   x() {return m_x;}
	void  setX(int x);
	void  setY(float min, float max);
	void  setYFactor();
	int	  xAt(int i);
	int	  xWOne();

protected:
	void resizeEvent(QResizeEvent *e);

public slots:
	void OnXYChanged();

private:
	void setupInfoBox();

private:
	//Ui::XY ui;

	float m_ymin, m_ymax;
	int m_x;
	float m_yfactor;

	//info
	QLabel *m_lb_size;
	QLabel *m_lb_x;
	QLabel *m_lb_y;
	QPushButton *btn_start;
	QLineEdit *m_le_x, *m_le_ymin, *m_le_ymax;
};

#endif // XY_H
