var QuizQuestions = [{"text":"Dual Photography Is:","options":[{"text":"Controlling light sources allowing more data than traditional photography","correct":true},{"text":"The use of multiple cameras to produce a novel image","correct":false},{"text":"Doubling the number of lenses to meet light source","correct":false}]},{"text":"Dual photography mostly came from this university","options":[{"text":"Harvard","correct":false},{"text":"Georgia Tech","correct":false},{"text":"Stanford","correct":true}]},{"text":"Which surface type is a diffuse surface? ","options":[{"text":"Specular","correct":false},{"text":"Mirrored","correct":false},{"text":"Matte","correct":true}]},{"text":"What are the elements of computational photography: ","options":[{"text":"Illumination, Import, Modification Sharing","correct":false},{"text":"Illumination, Optics, Sensor, Processing, Display","correct":true},{"text":"Capturing, Importing, Editing, Printing","correct":true}]},{"text":"What are the parts of a novel camera:","options":[{"text":"Generalized optics, Aperture, Sensor, Processing","correct":true},{"text":"Sensor, Lens, Filters, Functions","correct":false},{"text":"Camera, Software","correct":false}]}];

const InitTime = 30;
var Timer;
var CorrectCount = 0;
var IncorrectCount = 0;
var Incorrect = [];
var idx = 0;
var Answers = [];
var time = InitTime;
var questionTemplate = "<div class='adept-quiz-timer'></div><div><div style='align:right;color:#333;text-align:right'>Question {{index}} of {{length}}</div><h3>{{{text}}}</h3></div><ul class='adept-questions'>{{#options}}<li><input name='foo' type='radio' data-correct=\"{{#correct}}1{{/correct}}{{^correct}}0{{/correct}}\" /> {{{text}}}</li>{{/options}}</ul><div style='padding-top:20px;'><a href='#' onclick='EvalAnswer();' class='adept-btn big'>Next</a></div>";

var Audio = {
	correct: new Audio('/assets/sound/adept_correct.mp3'),
	incorrect: new Audio('/assets/sound/adept_incorrect.mp3'),
	success: new Audio('/assets/sound/adept_success.mp3'),
	failure: new Audio('/assets/sound/adept_failure.mp3')
}

// http://stackoverflow.com/questions/2450954/how-to-randomize-shuffle-a-javascript-array
function shuffle(array) {
  var currentIndex = array.length, temporaryValue, randomIndex;

  // While there remain elements to shuffle...
  while (0 !== currentIndex) {

    // Pick a remaining element...
    randomIndex = Math.floor(Math.random() * currentIndex);
    currentIndex -= 1;

    // And swap it with the current element.
    temporaryValue = array[currentIndex];
    array[currentIndex] = array[randomIndex];
    array[randomIndex] = temporaryValue;
  }

  return array;
}

function UpdateQuiz() {
	//console.log(QuizQuestions,Answers);
	$('#adept-status').html(Mustache.render("{{#questions}}<div class='adept-answer {{#correct}}correct{{/correct}}'></div>{{/questions}}",{ questions: Answers}));
	Q = QuizQuestions[idx];
	$('#adept-quiz').html(Mustache.render(questionTemplate,Q));
	idx++;
	StartQuizTimer();
}

function countdown() {

		time--;
			$('.adept-quiz-timer').text(time);
		if (time > 0) {
			Timer = setTimeout(countdown,1000);
		} else {
			EvalAnswer();
		}

}

function StartQuizTimer() {

	$('.adept-quiz-timer').text(time);
	Timer = setTimeout(countdown, 1000);
}

function EvalAnswer() {

	isLast = false;
	if (parseInt(idx) > parseInt(QuizQuestions.length-1)) {
		isLast = true;
	}

	s = $(':checked').attr('data-correct');

	if (s == 1) {
		Answers.push({correct:true});
		CorrectCount++;
		if (!isLast) {
			Audio.correct.volume = 0.1;
			Audio.correct.play();
		}
	} else {
		Answers.push({correct:false});		
		IncorrectCount++;
		Incorrect.push($(this).closest('h3'));
		if (!isLast) {
			Audio.incorrect.play();	
			Audio.incorrect.volume = 0.7;					
		}
	}
		myChart.data.datasets[0].data[1] = CorrectCount;
		myChart.data.datasets[0].data[0] = IncorrectCount;
		myChart.update();	
	clearTimeout(Timer);
	time = InitTime;


	if (parseInt(idx) > parseInt(QuizQuestions.length-1)) {
		TallyQuiz();
		return false;
	} else {

	}


	UpdateQuiz();
}

function TallyQuiz() {
	total = 0;
	correct = 0;
	Answers.forEach(function(a) {
		total++;
		if (a.correct) {
			correct++;
		}
	})
	avg = correct/total;
	if (avg > .5) {
		$('.adept-quiz-container').html(Mustache.render("<div><h2>Congratulations</h2>  You passed Level 1 of Learning Javascript with a score of "+(avg*100)+"%.<p><a href='/dashboard'>Back to your dashboard</a>",{ questions: Answers}));		
		Audio.success.volume = .05;
		Audio.success.play();
	} else {
		$('.adept-quiz-container').html(Mustache.render("Sorry, you did not pass this quiz.  <a href='/course/?/quizzes/?' class='btn'>Try Again</a>",{ questions: Answers}));
			Audio.failure.volume = .03;		
		Audio.failure.play();				
	}
}

function StartQuiz() {


	$.getJSON("/api/quiz/"+Quiz,function(d) {
		q = [];
		i = 1;

		$(d.Questions).each(function() {
			this.answers = shuffle(this.answers);
			q.push({index:i,length: d.Questions.length, text:this.text, options: this.answers});
			i++;
		});

		QuizQuestions = q;
		UpdateQuiz();		
	});


}

$(document).ready(function() {

	StartQuiz();

})