//类继承中的重名成员
#include<iostream>
using namespace std;

/*
自己猜想：

对于子类中的与父类重名的成员，c++编译器会单独为子类的这个成员变量再开辟一块内存空间，
把这个重名的成员变量当成子类的独有属性
在子类对象中如果访问重名成员，会默认访问子类独有那个成员变量，而不是访问父类的成员变量

对于从父类继承的所有成员，c++编译器会在子类对象的内存空间前部分放置父类的所有成员
父类的成员函数访问的都是  子类对象中  父类内存空间那部分的成员变量，
--如果不是这样，那么父类中能访问子类成员变量？
又因为子类对象的重名成员  是独立于父类    属于子类对象的一个成员变量

对于被子类继承的父类的成员函数 访问的是父类中的那个重名成员的解释
根据类的对象管理模型，类的非静态成员函数就是一个全局函数，只是函数中默认有一个对象指针参数
对于父类成员函数而言，父类成员函数有一个参数就是父类对象指针，
此时子类对象调用父类成员函数，不就是类的赋值兼容性原则吗--父类指针调用子类对象，
父类对象指针不强转成子类对象指针，那么这个父类指针不是只能访问父类对象的成员吗？
因此得出结论，子类对象调用继承于父类的成员函数 访问的是父类中的那个重名成员
*/

/*
总结：
派生类定义了与基类同名的成员，在派生类中访问同名成员时屏蔽了基类的同名成员
在派生类中使用基类的同名成员，显式的使用类名限定符 -----    类名::成员
*/

class PointA{
public:
    int x;
    int y;
    void PrintA(){
        cout << "x=" << x << ";y=" << y << endl;
    }
    void PrintB(){
        cout << "x=" << x << ";y=" << y << endl;
    }
};

class PointB :public PointA{
public:
    PointB(){
        x = 3;
        y = 8;
        z = 7;
        //在子类内部中给可访问父类成员属性赋值
        PointA::y = 11;
    }
    int y;
    int z;
    void Test1(){
        cout << "z=" << z << ";y=" << y << endl;
    }
    void PrintB(){
        cout << "z=" << z << ";y=" << y << endl;
    }
};

void ProtectA(){
    PointB pb1;
    pb1.PrintA();
    cout << "----------------" << endl;
    pb1.Test1();
    cout << "----------------" << endl;
    pb1.y = 33;
    //在子类外部中给可访问父类成员属性赋值
    pb1.PointA::y = 44;
    pb1.PrintA();
    cout << "----------------" << endl;
    pb1.Test1();
}

void ProtectB(){
    PointB pb1;
    //对于重名函数调用规则和重名变量规则类似
    pb1.PrintB();
    cout << "--------------------" << endl;
    pb1.PointA::PrintB();
}

class X {
public:
	int m;
};
class Y : virtual public X {
public:
	int m;
};
int main(){
    //ProtectA();
    //ProtectB();
	Y y;
	y.m = 10;
	y.X::m = 999;
	std::cout << "y.m = " << y.m << "\n";
	std::cout << "y.X::m = " << y.X::m << "\n";
}
