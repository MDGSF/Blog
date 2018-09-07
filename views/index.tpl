
<ul class="layui-timeline">

  {{with .Posts}}
  {{range .}}
  <li class="layui-timeline-item">
    <i class="layui-icon layui-timeline-axis">&#xe63f;</i>
    <div class="layui-timeline-content layui-text">
      <h3 class="layui-timeline-title"><a href="{{.ReqURL}}" rel="bookmark">{{.Title}}</a></h3>
      <ul>
        <li> {{.Time}} </li>
        <li> {{.Author}} </li>
        <li> {{.Tags}} </li>
      </ul>
    </div>
  </li>
  {{end}}
  {{end}}

</ul>