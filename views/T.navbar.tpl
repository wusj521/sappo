{{define "navbar"}}
  <div>
  		<!--<a class="navbar-brand" href="/">我的博客</a>-->
    		<ul class="nav nav-pills">
  				<li {{if .IsShouy}} class="active"{{end}}><a href="/">首页</a></li>
          <!--<li {{if .IsCategory}} class="active"{{end}}><a href="/category">分类</a></li>-->
					<li {{if .IsDaisp}} class="active"{{end}}><a href="/daisp">待审批</a></li>
					<li {{if .IsSappo}} class="active"{{end}}><a href="/sappo">下载PO单</a></li>
  				{{if .IsLogin}}
  				<li class="pull-right"><a href="/login?exit=true">退出</a></li>
  				{{else}}
  				<li class="pull-right"><a href="/login">登录</a></li>
  				{{end}}
			</ul>
		</div>
{{end}}