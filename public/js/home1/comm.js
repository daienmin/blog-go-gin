$(document).ready(function () {

	/*var obj;
	var As=document.getElementById('nav').getElementsByTagName('a');
	obj = As[0];
	for(var i=1;i<As.length;i++){
		if(window.location.href.indexOf(As[i].href)>=0)
			obj=As[i];
	}
	obj.id='selected';*/

	$("#starlist li:nth-child(1) span").click(function(){
		$("#starlist li:nth-child(1) ul.sub").toggleClass("open");
	});
	$("#starlist li:nth-child(2) span").click(function(){
		$("#starlist li:nth-child(2) ul.sub").toggleClass("open");
	});
	$("#starlist li:nth-child(3) span").click(function(){
		$("#starlist li:nth-child(3) ul.sub").toggleClass("open");
	});
  
	//nav
	$("#mnavh").click(function(){
    $("#nav").toggle();
	$("#mnavh").toggleClass("open");
	});
  
  	//scroll to top
    var offset = 300,
        offset_opacity = 1200,
        scroll_top_duration = 700,
        $back_to_top = $('.cd-top');

    $(window).scroll(function () {
        ($(this).scrollTop() > offset) ? $back_to_top.addClass('cd-is-visible') : $back_to_top.removeClass('cd-is-visible cd-fade-out');
        if ($(this).scrollTop() > offset_opacity) {
            $back_to_top.addClass('cd-fade-out');
        }
    });
    $back_to_top.on('click', function (event) {
        event.preventDefault();
        $('body,html').animate({
                scrollTop: 0,
            }, scroll_top_duration
        );
    });
			
});