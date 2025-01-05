// static/js/main.js
// Audio Player Script
document.getElementById('playAudio')?.addEventListener('click', () => {
    const audio = document.getElementById('welcomeAudio');
    if (audio.paused) {
        audio.play();
    } else {
        audio.pause();
    }
});

// Quiz Score Handler
document.body.addEventListener('htmx:afterOnLoad', function(evt) {
    const triggeredEvents = evt.detail.trigger?.getAttribute('hx-trigger');
    if (triggeredEvents && triggeredEvents.includes('updateScore')) {
        const score = document.getElementById('currentScore');
        const attempted = document.getElementById('questionsAttempted');
        const correct = document.getElementById('correctAnswers');

        attempted.textContent = parseInt(attempted.textContent) + 1;

        if (evt.detail.isCorrect) {
            correct.textContent = parseInt(correct.textContent) + 1;
            score.textContent = parseInt(score.textContent) + 10;
        }
    }
});