<!DOCTYPE html>
<html>
<head>
  {{ .HTMLHead }}
</head>
<body>

<div class="container">
  {{.LayoutContent}}
</div>

<div>
  {{.SideBar}}
</div>

{{.Scripts}}

</body>
</html>