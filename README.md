**项目概览：**

![压测平台功能][1]

---

**项目介绍：**

1. 在ngrinder基础上做的二次开发
2. 完整功能由ngrinder、ngrinder-sampling组成，同时需要安装agent、monitor
![组成][2]
3. ngrinder-sampling负责场景的创建、ngrinder负责下发测试和报告收集、agent服务脚本测试、monitor负责代理机器的硬件使用情况
4. 项目使用mysql做数据存储，场景生成脚本使用svn做版本控制，脚本由groovy生成

---

**项目安装:**

一. ngrinder
    
    所需环境:jdk1.8、maven、idea
    
    1. 下载 http://192.168.0.167/changli/ngrinder.git
    2.  导入idea后，在父类的pom.xml文件中添加repository（因为有几个包有问题，没有发布在中央仓库）
        
        ```
        <repositories>
            <repository>
        			<id>ngrinder-core</id>
        			<url>https://github.com/nhnopensource/nhnopensource.maven.repo/raw/master/releases</url>
        	</repository>
        </repositories>
        ```
        ps：如果下载有问题，可以挂上vpn进行操作。
    3. 修改项目的jdk版本，设置为1.8.X
    4. 将 ngrinder-controller打成的war包上传至服务器
    5. 启动
    
    ```
    nohup java -XX:MaxMetaspaceSize=512m -jar /root/ngrinder/ngrinder-controller-3.4.3.war --port 8080 < /dev/null &
    ```
    6. 启动后被解压到~/.ngrinder/下， 日志查看等可到该处查看
    7. 可在此时修改database.conf、system.conf系统配置、数据库链接
        注意服务器要在/etc/hosts中添加ip的映射，不然会存在问题
二. ngrinder-sampling
    
    所需环境:golang1.11.13、beego、bee
    
    1. 下载 http://192.168.0.167/changli/ngrinder-sampling.git
    2. 执行根目录下的pack.bat打包出ngrinder-sampling.tar.gz上传至服务器
    3. 解压后可修改conf/app.conf文件，修改数据库、系统配置
    4. 执行nuhup ./ngrinder-sampling & 启动
三. agent
    
    1. 登陆http://192.168.0.11:8080/
    2. 下载代理，获取到ngrinder-agent-3.4.3-192.168.0.11.tar
    3.  上传到代理机器，解压 执行run_agent_bg.sh
   
   ![下载代理][3] 
四. monitor
    
    1. 下载监控ngrinder-monitor-3.4.3.tar
    2. 上传至服务器，解压 执行run_monitor_bg.sh

---

**项目讲解：**

1. 登陆 http://192.168.0.11:8090/
    
    钉钉扫码登陆或者账号登陆，账号只能由管理员在ngrinder后台创建
    ![登陆][4]
2. 公告
    
    公告可由管理员在ngrinder进行修改，及时通知使用用户是否有动态
    ![公告][5]
3. 创建压测
    
   1. 场景配置：配置接口基本信息
       ![场景配置][6]
       body参数可定义为${**}，执行时随机获取数据源内的数据
   2. 数据源配置：压测参数来源，csv格式,记得点击保存、否则不生效
       ![数据源][7]
   3. 施压配置
       ![施压配置][8]
   4. 高级设置
       ![高级设置][9]
   5. 添加监控
        ![添加监控][10]
   6. 保存配置
   
        将场景压测保存，并跳转到压测场景列表
   7. 调试场景
   
        对场景压测进行单次测试，校验录入执行是否符合结果预期.
        返回结果为本场景请求信息
        ![调试场景][11]
   8. 保存去压测
        
         分为马上运行和预约两部分，预约需要制定执行时间.点击后跳转到压测场景列表    
         ![保存去压测][12]
4. 压测场景
        
      创建的场景在此查看，可搜索、修改、删除
      ![压测场景][13]                    
5. 压测报告
    
      所有进行的压测可在此进行查看、删除
      ![压测报告][14]
      
      1. 查看详情
         
         可查看压测的详细信息，生成的图标，以及检测机器的使用情况。也可下载压测报告的csv文件
         ![查看详情][15]
      2.  查看采样日志
         
          采样率按照10%的比例进行的收集
          ![采样列表][16]
          
          ![采样详情][17]
      3. 删除
        
          删除后，本次压测对应的数据、压测结果等全部删除不可找回。删除时，请仔细考虑。                       
        
[1]: views/images/ngrinder_sampling.png
[2]: views/images/ngrinder_struct.png
[3]: views/images/ngrinder_agent_download.png
[4]: views/images/ngrinder_login.png
[5]: views/images/ngrinder_announcement.png
[6]: views/images/ngrinder_scenes_create.png
[7]: views/images/ngrinder_data_file.png
[8]: views/images/ngrinder_test_config.png
[9]: views/images/ngrinder_test_advanced.png
[10]: views/images/ngrinder_monitor.png
[11]: views/images/ngrinder_test_rst.png
[12]: views/images/ngrinder_test_schedule.png
[13]: views/images/ngrinder_scenes_list.png
[14]: views/images/ngrinder_report_list.png
[15]: views/images/mgrinder_report_detail.png
[16]: views/images/ngrinder_sampling_list.png
[17]: views/images/ngrinder_sampling_detail.png