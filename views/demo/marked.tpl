<!doctype html>
<html>
<head>
  <meta charset="utf-8"/>
  <title>Marked in the browser</title>
</head>
<body>
  <div id="content"></div>
  <script src="https://cdn.jsdelivr.net/npm/marked/marked.min.js"></script>
  <!-- <script src="/static/marked/marked.min.js"></script> -->
  <script>
    document.getElementById('content').innerHTML =
      marked('{{.Post}}');
  </script>
</body>
</html>