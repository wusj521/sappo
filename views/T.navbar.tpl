{{define "navbar"}}
  <div>
  		<!--<a class="navbar-brand" href="/">我的博客</a>-->
    		<ul class="nav nav-pills">
  				<li {{if .IsShouy}} class="active"{{end}}><a href="/">首页</a></li>
          <!--<li {{if .IsCategory}} class="active"{{end}}><a href="/category">分类</a></li>-->
					<li class="dropdown" >
            <a href="#" class="dropdown-toggle" data-toggle="dropdown">
              待审
              <b class="caret"></b>
            </a>
					  <ul class="dropdown-menu">
					     <li {{if .IsDaisp}} class="active"{{end}}><a href="/daisp">待审PO</a></li>
				       <li {{if .IsDaisppr}} class="active"{{end}}><a href="/daisppr">待审PR</a></li>
					  </ul>
					</li>

					<li {{if .IsSappo}} class="active"{{end}}><a href="/sappo">下载SAP单据</a></li>
  				{{if .IsLogin}}
  				<li class="pull-right"><a href="/login?exit=true">退出</a></li>
  				{{else}}
  				<li class="pull-right"><a href="/login">登录</a></li>
  				{{end}}
			</ul>
		</div>
{{end}}