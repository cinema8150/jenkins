TODO: 
【】 默认分支
【】 仅在变更时编译（距离上次编译成功有变动，需排查编译失败问题）


"java -jar jenkins-cli.jar -s " + viper.GetString("jenkins.url") + " -webSocket -auth " + viper.GetString("jenkins.auth") + " build test_iOS_module_deploy_to_project -p project=" + project + " -p branch=" + branch + " -p modules_info=" + mdList




java -jar jenkins-cli.jar -s http://app.jenkins.jlgltech.com:5001 -webSocket -auth ekg_yin1:110c6bcec70f8ddcf51610f6fac1755a8d build "iOS_Cybertron_intl" -p platform=pgy -p configuration=AdHoc -p export_method=ad-hoc -p branch=origin/develop/1.6.0

Neither -s nor the JENKINS_URL env var is specified.
Jenkins CLI
Usage: java -jar jenkins-cli.jar [-s URL] command [opts...] args...
Options:
 -s URL              : the server URL (defaults to the JENKINS_URL env var)
 -http               : use a plain CLI protocol over HTTP(S) (the default; mutually exclusive with -ssh)
 -webSocket          : like -http but using WebSocket (works better with most reverse proxies)
 -ssh                : use SSH protocol (requires -user; SSH port must be open on server, and user must have registered a public key)
 -i KEY              : SSH private key file used for authentication (for use with -ssh)
 -noCertificateCheck : bypass HTTPS certificate check entirely. Use with caution
 -noKeyAuth          : don't try to load the SSH authentication private key. Conflicts with -i
 -user               : specify user (for use with -ssh)
 -strictHostKey      : request strict host key checking (for use with -ssh)
 -logger FINE        : enable detailed logging from the client
 -auth [ USER:SECRET | @FILE ] : specify username and either password or API token (or load from them both from a file);
                                 for use with -http.
                                 Passing credentials by file is recommended.
                                 See https://www.jenkins.io/redirect/cli-http-connection-mode for more info and options.
 -bearer [ TOKEN | @FILE ]     : specify authentication using a bearer token (or load the token from file);
                                 for use with -http. Mutually exclusive with -auth.
                                 Passing credentials by file is recommended.

java -jar jenkins-cli.jar -s http://app.jenkins.jlgltech.com:5001 -ssh build "iOS_Cybertron_intl" -p platform=pgy -p configuration=AdHoc -p export_method=ad-hoc -p branch=origin/develop/1.6.0

java -jar jenkins-cli.jar -s http://app.jenkins.jlgltech.com:5001 -webSocket -ssh -user ekg_yin1 build "iOS_Cybertron_intl" -p platform=pgy -p configuration=AdHoc -p export_method=ad-hoc -p branch=origin/develop/1.6.0