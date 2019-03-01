---
title: "Remote"
description: "Java 远程调试"
date: 2019-02-14T17:19:30+08:00
draft: true
toc: true
---

# JPDA
JPDA（Java Platform Debugger Architecture）定义了一套完整独立的体系，它由三个相对独立的层次共同组成，而且规定了它们三者之间的交互方式，或者说定了通讯接口。

JVM自身支持了远程调试，而我们常用的开发工具Eclipse也提供了远程调试的功能。

例如，在tomcat中的jpda端口是8000，tomcat以jpda的方式运行起来后，就可以通过Eclipse连接上。

如下图所示，在1处添加一个Java Remote Application配置；在2处选择要调试的工程；在3处选择调试模式——Attach的意思就是说Eclipse要主动连接远程的Java程序；4处的地址和端口就不解释了。

{{< img "/images/java/debug_remote_java_app-300x269.png" >}}

在下图所示的选项卡中可以设置当前调试中的源码，通过2处的按钮可以浏览；勾选3处的复选框之后，会根据上面列表中的所有位置查找源码，如果找到多份的话弹出选中对话框。

{{< img "/images/java/debug_remote_java_app_source.png" >}}

Eclipse提供了多种源码选择的方式，通过下图可以看到：你可以把源码作为Java工程来选择，也可以作为普通的文件来选择。

{{< img "/images/java/add_source-300x199.png" >}}

# JVMTI
JVMTI（Java Virtual Machine Tool Interface）就是Java虚拟机工具接口。

这里有一个开源的实现：http://harmony.apache.org/subcomponents/drlvm/index.html

# JDWP
JDWP（Java Debug Wire Protocol）是为Java调试而设计的一个通讯交互协议。调试既可以在同一台机器上，也可以使远程的。JDK包含一个默认的实现jdwp.dll。JVM允许灵活地使用其他协议来替代JDWP。

你可以使用以下的参数以调试模式启动 JVM：

-Xdebug -Xnoagent -Xrunjdwp:transport=dt_socket,address=8000,server=y,suspend=n

-Xrunjdwp     JVM 加载 jdwp.dll

transport=dt_socket   使用 Socket 传输

address      表示调试端口

server=y     表示 JVM 作为服务器，建立 Socket

suspend=n    表示启动过程中，JVM 不会挂起去等待调试客户端连接

# JDI
JDI（Java Debug Interface）是三个模块中最高层的接口。这是一组Java接口。

Eclipse默认提供的调试客户端是org.eclipse.jdt.debug.