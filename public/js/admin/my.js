function succ(title,msg,url){
	$.gritter.add({
		title:	title,
		text:	msg,
		time: 200,  
		speed:1000,
		sticky: false,
		after_close: function(){
			window.location.href=url;
		}
	});
}

function err(msg){
		$.gritter.add({
		title:	"错误!",
		text:	msg,
		time: 1000,  
		speed:1000,
		sticky: false,
	});
}

$(function () {
	var url = location.href;
	$("#sidebar li.submenu").each(function (i,v) {
		var control = $(v).attr("data-control");
		if (url.indexOf(control) > -1) {
			$(v).addClass('open').find('ul').show();
			return false;
		}
    });
});
	