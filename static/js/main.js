$(window).load(onLoad);

function onLoad() {
	$('#firstname-container').hide();
	$('#guesser').submit(function(e) {
		console.log("SUBMIT!");
		$('#firstname-container').hide();
		$.ajax({
			type: "POST",
			url: "/name",
			data: $("#guesser").serialize(),
			success: function(data) {
				console.log(data);
				j = JSON.parse(data);
				$('#firstname').text(j.first_name);
				$('#firstname-container').show();
			}
		});
		return false;
	});
}