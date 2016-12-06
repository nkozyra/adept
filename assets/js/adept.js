/*	Adept.js */

$(document).ready(function() {


	$('.adept-header-user').click(function() {
		$(this).toggleClass('active');
		$('.adept-header-user-opts').toggle();
	});

});