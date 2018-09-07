

<div class="layui-container">  
    <div class="layui-row">
      <div class="layui-col-md4">
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
      </div>
    </div>
    <div class="layui-row">
      <div id="test1"></div>
    </div>
</div>

<script src="/static/jquery/jquery-1.12.4/jquery.min.js"></script>
<script src="/static/layui-v2.4.3/layui/layui.js"></script>
<script>
layui.use('laypage', function(){
  var laypage = layui.laypage;
  
  //执行一个laypage实例
  laypage.render({
    elem: 'test1' //注意，这里的 test1 是 ID，不用加 # 号
    ,count: {{.PageCount}} //数据总数，从服务端得到
    ,limit: {{.PageLimit}} //每页显示多少条
    ,limits: [10, 50, 100, 200, 300, 400, 500, 1000]
    ,layout: ['limit', 'prev', 'page', 'next']
    ,jump: function(obj, first){
      //obj包含了当前分页的所有参数，比如：
      console.log(obj.count);
      console.log(obj.curr); //得到当前页，以便向服务端请求对应页的数据。
      console.log(obj.limit); //得到每页显示的条数
      
      //首次不执行
      if(!first){
        //do something
        $.get("/pg?pageCount="+obj.count+"&pageLimit="+obj.limit+"&curPage="+obj.curr)
      }
    }
  });
});
</script>