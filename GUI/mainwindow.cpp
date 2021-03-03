#include "mainwindow.h"
#include "ui_mainwindow.h"
#include "xy.h"

MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
    , ui(new Ui::MainWindow)
{
    ui->setupUi(this);

	setCentralWidget(new XY(this));
}

MainWindow::~MainWindow()
{
    delete ui;
}

