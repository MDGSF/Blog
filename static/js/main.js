(function($){

    // change locale and reload page
	$(document).on('click', '.lang-changed', function(){
		var $e = $(this);
		var lang = $e.data('lang');
		Cookies.set('lang', lang, {path: '/', expires: 365});
		window.location.reload();
	});

})(jQuery);