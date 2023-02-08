<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta title="Veinmind Report">
    <link rel="stylesheet" href="https://cdn.dvkunion.cn/veinmind/semantic.min.css">
    <link rel="stylesheet" href="https://cdn.dvkunion.cn/veinmind/fonts/icon.min.css">
    <style>
        .body {
            background: RGB(234, 234, 239);
        }

        .ui.inverted.menu {
            background-color: #310056;
            font-size: 1rem;
        }

        .ui.fixed.menu {
            width: 100%;
        }

        .ui.menu .item > img:not(.ui) {
            font-size: 20px;
            width: 122px;
        }

        .masthead.segment.bg1 .container {
            margin-left: auto !important;
            margin-right: auto !important;
            min-height: 400px;
            -webkit-box-align: center;
            -webkit-align-items: center;
            -ms-flex-align: center;
            align-items: center;
            display: -webkit-box;
            display: -webkit-flex;
            display: -ms-flexbox;
            display: flex;
            -webkit-box-pack: center;
            -webkit-justify-content: center;
            -ms-flex-pack: center;
            justify-content: center;
        }

        .masthead .introduction {
            position: relative;
            clear: both;
            display: block;
            text-align: center;
        }

        .masthead.segment h1 > .library {
            display: block;
            font-size: 1.25em;
            margin-bottom: 10px;
            font-weight: bold;
        }

        .masthead.segment h1 .tagline {
            font-size: 0.75em;
        }

        .ui.version.label {
            background-color: transparent;
            font-weight: 16;
            width: 128px;
        }

        .part {
            margin: 2em 0em;
            position: relative;
            -webkit-tap-highlight-color: transparent;
        }

        .ui.header > .icon {
            display: inline;
            font-size: 1em;
            margin-right: 10px;
        }

        .part i.code {
            cursor: pointer;
            position: absolute;
            top: 2rem;
            right: 0rem;
            margin: 0;
            opacity: 0.5;
            font-size: 15px;
            color: #000000;
            -webkit-transition: opacity 0.3s ease-out;
            -moz-transition: opacity 0.3s ease-out;
            -o-transition: opacity 0.3s ease-out;
            -ms-transition: opacity 0.3s ease-out;
            transition: opacity 0.3s ease-out;
        }

        .part i.code, #part h4 + .part i.code, #part h4 + .part i.code {
            top: 0em;
        }
    </style>
</head>
<body>
<div class="body">
    <div class="ui fixed inverted main menu">
        <div class="ui container">
            <a href="https://veinmind.chaitin.com/docs/" class="launch icon item" target="_blank">
                <img class="logo" src="https://cdn.dvkunion.cn/veinmind/logo_white.png"/>
            </a>
            <div class="ui inverted right labeled icon menu">
                <a class="item" href="https://veinmind.chaitin.com/" target="_blank">
                    <i class="home icon"></i>
                    HomePage
                </a>
                <a class="item" href="https://github.com/chaitin/veinmind-tools" target="_blank">
                    <i class="github icon"></i>
                    Github
                </a>
                <a class="item" href="https://veinmind.chaitin.com/docs/" target="_blank">
                    <i class="file alternate outline icon"></i>
                    Document
                </a>
                <a class="item" href="https://veinmind.chaitin.com/static/images/QRcode.8ffb1a6.png" target="_blank">
                    <i class="wechat icon"></i>
                    Contact Us
                </a>
            </div>
        </div>
    </div>
    <div class="masthead segment bg1">
        <div class="ui container">
            <div class="introduction">
                <img class="ui version label" src="https://cdn.dvkunion.cn/veinmind/favicon.svg"/>
                <h1 class="ui header">
                    <span class="library">
                      Veinmind Tools
                    </span>
                    <span class="tagline">
                        容器安全见筋脉，望闻问切治病害
                    </span>
                </h1>
            </div>
        </div>
    </div>

    <div class="ui container">
        {{- if or .BasicImage .BasicContainer .BasicCluster .Asset}}
        <div class="ui active intro tab" data-tab="overview">
            <h2 class="ui dividing center header"><i class="search icon"></i>资产详情
                <a class="anchor"
                   id="detail"></a></h2>
            {{- if .BasicImage }}
            <div class="highlighted part" data-class="BasicImage">
                <h4 class="ui header">镜像信息</h4>
                <i class="fitted icon code"></i>
                <a class="anchor" id="BasicImage"></a>
                {{ .BasicImage }}
                <div class="ui ignored hidden divider"></div>
            </div>
            {{- end }}
            {{- if .BasicContainer }}
            <div class="highlighted part" data-class="BasicContainer">
                <h4 class="ui header">容器信息</h4>
                <i class="fitted icon code"></i>
                <a class="anchor" id="BasicContainer"></a>
                {{ .BasicContainer }}
                <div class="ui ignored hidden divider"></div>
            </div>
            {{- end }}
            {{- if .BasicCluster }}
            <div class="highlighted part" data-class="BasicCluster">
                <h4 class="ui header">集群信息</h4><i class="fitted icon code"></i>
                <a class="anchor" id="BasicCluster"></a>
                {{ .BasicCluster }}
                <div class="ui ignored hidden divider"></div>
            </div>
            {{- end }}
            {{- if .Asset }}
            <div class="highlighted part" data-class="Asset">
                <h4 class="ui header">应用软件信息</h4><i class="fitted icon code"></i>
                <a class="anchor" id="Asset"></a>
                {{ .Asset }}
                <div class="ui ignored hidden divider"></div>
            </div>
            {{- end }}
        </div>
        <div class="ui ignored hidden divider"></div>
        {{- end }}
        {{- if or .Vulnerability .MaliciousFile .SensitiveFile .SensitiveEnv .SensitiveHistory .Weakpass .Backdoor .AbnormalHistory .IaC .Webshell .Escape .UnsafeMount}}
        <div class="ui active intro tab" data-tab="overview">
            <h2 class="ui dividing center header"><i class="ambulance icon"></i>风险详情<a class="anchor"
                                                                                       id="risk"></a></h2>
            {{- if .Vulnerability }}
            <div class="highlighted part" data-class="bugs">
                <h4 class="ui header">应用漏洞</h4>
                <i class="fitted icon code"></i>
                <a class="anchor" id="bugs"></a>
                {{ .Vulnerability }}
                <div class="ui ignored hidden divider"></div>
            </div>
            {{- end }}
            {{- if .MaliciousFile }}
            <div class="highlighted part" data-class="MaliciousFile">
                <h4 class="ui header">恶意文件</h4>
                <i class="fitted icon code"></i>
                <a class="anchor" id="MaliciousFile"></a>
                {{ .MaliciousFile }}
                <div class="ui ignored hidden divider"></div>
            </div>
            {{- end }}
            {{- if or .SensitiveFile .SensitiveEnv .SensitiveHistory }}
            <div class="highlighted part" data-class="Sensitive">
                <h4 class="ui header">敏感信息</h4><i class="fitted icon code"></i>
                <a class="anchor" id="Sensitive"></a>
                {{ .SensitiveFile }}
                {{ .SensitiveEnv }}
                {{ .SensitiveHistory }}
                <div class="ui ignored hidden divider"></div>
            </div>
            {{- end }}
            {{- if .Weakpass }}
            <div class="highlighted part" data-class="Weakpass">
                <h4 class="ui header">弱口令</h4><i class="fitted icon code"></i>
                <a class="anchor" id="Weakpass"></a>
                {{ .Weakpass }}
                <div class="ui ignored hidden divider"></div>
            </div>
            {{- end }}
            {{- if .Backdoor }}
            <div class="highlighted part" data-class="Backdoor">
                <h4 class="ui header">后门文件</h4><i class="fitted icon code"></i>
                <a class="anchor" id="Backdoor"></a>
                {{ .Backdoor }}
                <div class="ui ignored hidden divider"></div>
            </div>
            {{- end }}
            {{- if .AbnormalHistory }}
            <div class="highlighted part" data-class="AbnormalHistory">
                <h4 class="ui header">异常历史命令</h4><i class="fitted icon code"></i>
                <a class="anchor" id="AbnormalHistory"></a>
                {{ .AbnormalHistory }}
                <div class="ui ignored hidden divider"></div>
            </div>
            {{- end }}
            {{- if .IaC }}
            <div class="highlighted part" data-class="IaC">
                <h4 class="ui header">错误IaC配置</h4><i class="fitted icon code"></i>
                <a class="anchor" id="IaC"></a>
                {{ .IaC }}
                <div class="ui ignored hidden divider"></div>
            </div>
            {{- end }}
            {{- if .Webshell }}
            <div class="highlighted part" data-class="Webshell">
                <h4 class="ui header">WebShell</h4><i class="fitted icon code"></i>
                <a class="anchor" id="Webshell"></a>
                {{ .Webshell }}
                <div class="ui ignored hidden divider"></div>
            </div>
            {{- end }}
            {{- if .Escape }}
            <div class="highlighted part" data-class="Escape">
                <h4 class="ui header">逃逸风险</h4><i class="fitted icon code"></i>
                <a class="anchor" id="Escape"></a>
                {{ .Escape }}
                <div class="ui ignored hidden divider"></div>
            </div>
            {{- end }}
            {{- if .UnsafeMount }}
            <div class="highlighted part" data-class="UnsafeMount">
                <h4 class="ui header">不安全的挂载</h4><i class="fitted icon code"></i>
                <a class="anchor" id="UnsafeMount"></a>
                {{ .UnsafeMount }}
                <div class="ui ignored hidden divider"></div>
            </div>
            {{- end }}
            {{- if .Others }}
            <div class="highlighted part" data-class="others">
                <h4 class="ui header">其他上报信息</h4><i class="fitted icon code"></i>
                <div class="ui ignored hidden divider"></div>
            </div>
            {{- end }}
        </div>
        {{- end }}
    </div>
</div>
</body>
<script src="https://cdn.dvkunion.cn/veinmind/jquery-3.1.1.min.js"></script>
<script src="https://cdn.dvkunion.cn/veinmind/semantic.min.js"></script>
<script>
    $('.menu .item')
        .tab();
    ;

    $('.ui.rating')
        .rating({
            initialRating: 3,
            maxRating: 5,
        })
        .rating('disable')
    ;

    $('#image-percent').progress();
</script>
</html>