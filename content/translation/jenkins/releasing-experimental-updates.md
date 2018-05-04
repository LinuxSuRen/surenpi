# 发布实验性版本的插件

[查看可用的开发者指导](https://jenkins.io/doc/developer/guides)

To simplify delivery of beta versions of plugins to interested users, the Jenkins project published an experimental update center. It will include alpha and beta versions of plugins, which aren’t usually included in the regular update sites.

# 创建试验性创建版本
Plugin releases that contain alpha or beta in their version number will only show up in the experimental update site. Note that it also serves regular releases, so the release of version 1.4 will result in 1.3-beta-2 no longer being available.

# 配置 Jenkins 使用试验性更新中心
Users who are interested in downloading experimental plugin releases can go to Plugin Manager, then to the Advanced tab, and configure the update center URL [https://updates.jenkins.io/experimental/update-center.json](https://updates.jenkins.io/experimental/update-center.json). Save, and then select Check Now. Experimental plugin updates will be marked as such on the Available and Updates tabs of the Plugin Manager.

Once you install the beta plugins that you wanted, you can switch back to the default [https://updates.jenkins.io/update-center.json](https://updates.jenkins.io/update-center.json) update center URL.

[original link](https://jenkins.io/doc/developer/publishing/releasing-experimental-updates/)
