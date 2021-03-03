#include "xy.h"
#include <QDebug>
#include <QFormLayout>

XY::XY(QWidget *parent)
	: QWidget(parent)
{
	//ui.setupUi(this);

	m_yfactor 	= 0;
	m_x 		= 0;
	m_ymin 		= 0;
	m_ymax		= 0;

	setupInfoBox();
	setStyleSheet("background-color:black;");
}

void XY::setupInfoBox()
{
	/* create debug labels */
	QWidget *info = new QWidget(this);
	QVBoxLayout *vlayout =  new QVBoxLayout;
	m_lb_x		= new QLabel(this);
	m_lb_y		= new QLabel(this);
	m_lb_size 	= new QLabel(this);
	m_le_x		= new QLineEdit(this);
	m_le_ymin	= new QLineEdit(this);
	m_le_ymax	= new QLineEdit(this);

	m_lb_y->setMinimumWidth(250);

	// formLayout
	QFormLayout *formlayout = new QFormLayout;
	formlayout->addRow(tr("Size:"), m_lb_size);
	formlayout->addRow(tr("X:"), m_lb_x);
	formlayout->addRow(tr("Y:"), m_lb_y);
	formlayout->addRow(tr("&X num:"), m_le_x);
	formlayout->addRow(tr("Y m&in:"), m_le_ymin);
	formlayout->addRow(tr("Y m&ax:"), m_le_ymax);
	m_le_x->setMaximumWidth(50);
	m_le_ymin->setMaximumWidth(50);
	m_le_ymax->setMaximumWidth(50);
	info->setLayout(formlayout);
	info->setWindowOpacity(0);
	info->setStyleSheet("background-color: rgba(0,0,0,0)");

	connect(m_le_x, SIGNAL(editingFinished()), this, SLOT(OnXYChanged()));
	connect(m_le_ymin, SIGNAL(editingFinished()), this, SLOT(OnXYChanged()));
	connect(m_le_ymax, SIGNAL(editingFinished()), this, SLOT(OnXYChanged()));
}

	//Events
void XY::resizeEvent(QResizeEvent *e)
{
	QString s = QString("size: %1, %2").arg(size().width()).arg(size().height());
	m_lb_size->setText(s);

	setX(m_x);
	setY(m_ymin, m_ymax);

	qDebug() << s;
}

XY::~XY()
{

}

// slots
void XY::OnXYChanged()
{
	int x = m_le_x->text().toInt();
	float ymin = m_le_ymin->text().toFloat();
	float ymax = m_le_ymax->text().toFloat();

	setX(x);
	setY(ymin, ymax);
	qDebug() << "xy changed";
}

void XY::setX(int x)
{
	m_x = x;

	QString s = QString("X: num: %1, w_one: %2").arg(x).arg(xWOne());
	m_lb_x->setText(s);
}
void  XY::setY(float min, float max) 
{
	m_ymin = min; 
	m_ymax = max;
	if (m_ymin > m_ymax)
	{
		float t = m_ymax;
		m_ymax = m_ymin;
		m_ymin = t;
	}
	setYFactor();

	QString s = QString("Y: (min,max):(%1, %2), yfactor: %3").arg(m_ymin).arg(m_ymax).arg(m_yfactor);
	m_lb_y->setText(s);
}

void XY::setYFactor()
{
	if (size().height() == 0)
		return;

	m_yfactor = (m_ymax - m_ymin) / size().height();
}

int XY::xAt(int i)
{
	if (i > m_x || i < 0)
		return 0;

	if (m_x == 0)
		return 0;

	int w_all, w_one;
	w_all = size().width();
	w_one = w_all / m_x;

	return i * w_one;
}

int XY::xWOne()
{
	return xAt(1) - xAt(0);
}
