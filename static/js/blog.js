
$("#topheader .nav a").on("click", function(){
   $("#topheader .nav").find(".active").removeClass("active");
   $(this).parent().addClass("active");
});

function validateSearchForm()
{
	var a=document.forms["searchForm"]["searchContent"].value.trim();
	if (a==null || a=="")
	{
		alert("Please Fill All Required Field");
		return false;
	}
}
