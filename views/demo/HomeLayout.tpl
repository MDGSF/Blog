<!DOCTYPE html>
<html>
<head>
  <title>MDGSF Blog</title>
  <meta charset="utf-8">

  <meta name="renderer" content="webkit"/>
  <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">

  <meta name="author" content="{{.Author}}">
  <meta name="keywords" content="mdgsf,{{.Website}}">
  <meta name="description" content="mdsg blog,html,css,js,golang">

  <link rel="stylesheet" href="/static/css/github-markdown.css">

  <!-- <link rel="stylesheet" href="/static/bootstrap-3.3.7-dist/css/bootstrap.min.css"> -->
  <link rel="stylesheet" href="/static/layui-v2.4.3/layui/css/layui.css">
</head>

<body class="layui-layout-body">
<div class="layui-layout">

<div class="layui-side layui-bg-molv">
  <div class="layui-side-scroll">
    <ul class="layui-nav layui-nav-tree layui-bg-molv layui-nav-side" lay-filter="test">
      <li class="layui-nav-item"><a href="">首页</a></li>
      <li class="layui-nav-item"><a href="">专题</a></li>
      <li class="layui-nav-item">
        <a class="" href="javascript:;">标签</a>
        <dl class="layui-nav-child">
          <dd><a href="javascript:;">列表一</a></dd>
          <dd><a href="javascript:;">列表二</a></dd>
          <dd><a href="javascript:;">列表三</a></dd>
          <dd><a href="">超链接</a></dd>
        </dl>
      </li>
      <li class="layui-nav-item"><a href="">归档</a></li>
      <li class="layui-nav-item"><a href="">链接</a></li>
      <li class="layui-nav-item"><a href="">关于</a></li>
    </ul>
  </div>
</div>

<div class="layui-body">
  <!-- 内容主体区域 -->
  <div style="padding: 15px;">
  <div class="container">
    {{.LayoutContent}}
  </div>
  </div>
</div>

<div class="layui-footer">
  <!-- 底部固定区域 -->
  © mdgsf.com - 底部固定区域
</div>

</div>

<!-- <script src="/static/jquery/jquery-1.12.4/jquery.min.js"></script> -->
<!-- <script src="/static/bootstrap-3.3.7-dist/js/bootstrap.min.js"></script> -->
<script src="/static/layui-v2.4.3/layui/layui.js"></script>
<script>
layui.use(['element'], function(){
		var element = layui.element;
});
</script> 

</body>
</html>