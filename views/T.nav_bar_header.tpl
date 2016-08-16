{{define "nav_bar_header"}}
<nav class="navbar navbar-static-top" role="navigation" style="margin-bottom: 0">
    <div class="navbar-header">
        <a class="navbar-minimalize minimalize-styl-2 btn btn-primary " href="#"><i class="fa fa-bars"></i> </a>
        <form role="search" class="navbar-form-custom" method="get" action="">
            <div class="form-group">
                <input type="text" placeholder="输入搜索..." class="form-control" name="search" id="top-search">
            </div>
        </form>
    </div>
    <ul class="nav navbar-top-links navbar-right">
        <li>
            <span class="m-r-sm text-muted welcome-message">欢迎使用听云NetopGO系统</span>
        </li>
        <!--
        {{if .IsViewOrder}}
        <li role="presentation"><a href="{{if .OrderFlag}}/workorder/my/list?pageAuth={{.PageAuth}}&pageDept={{.PageDept}}{{else}}/workorder/mydb/list?pageAuth={{.PageAuth}}&pageDept={{.PageDept}}{{end}}"><span style="color:blue;">待您完成工单</span><span class="badge" style="color:red;">{{.UnoverOrderNums}}</span></a></li>
        {{end}}
        -->
        {{if .IsViewOrder}}
        <li role="presentation"><a href="{{if .OrderFlag}}/workorder/my/list?pageAuth={{.PageAuth}}&pageDept={{.PageDept}}{{else}}/workorder/mydb/list?pageAuth={{.PageAuth}}&pageDept={{.PageDept}}{{end}}"><span style="color:blue;">待您完成工单</span><span class="badge" style="color:red;">{{.UnoverOrderNums}}</span></a></li>
        {{end}}
        <li>
            <a href="/logout">
                <i class="fa fa-sign-out"></i> Log out
            </a>
        </li>
    </ul>
</nav>
{{end}}
