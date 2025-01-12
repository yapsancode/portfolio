// Audio Player Script
document.getElementById('playAudio')?.addEventListener('click', () => {
    const audio = document.getElementById('welcomeAudio');
    if (audio.paused) {
        audio.play();
    } else {
        audio.pause();
    }
});

// Text Animation
const roles = [
    "Software Developer",
    "Software Engineer",
    "Full-Stack Developer",
    "Mobile App Developer",
];

const animatedText = document.getElementById('animatedText');

if (animatedText) {
    let currentRoleIndex = 0;
    let charIndex = 0;
    let isDeleting = false;

    function typeEffect() {
        const currentRole = roles[currentRoleIndex];
        
        // Fix: Ensure charIndex stays within bounds
        if (!isDeleting && charIndex > currentRole.length) {
            charIndex = currentRole.length;
        }
        if (isDeleting && charIndex < 0) {
            charIndex = 0;
        }

        const displayedText = currentRole.substring(0, charIndex);
        animatedText.textContent = displayedText;

        if (!isDeleting && charIndex >= currentRole.length) {
            // Reached the end of the word
            setTimeout(() => {
                isDeleting = true;
                typeEffect();
            }, 2000); // Pause at the end of the word
        } else if (isDeleting && charIndex === 0) {
            // Finished deleting
            isDeleting = false;
            currentRoleIndex = (currentRoleIndex + 1) % roles.length;
            setTimeout(typeEffect, 100);
        } else {
            // Still typing or deleting
            if (isDeleting) {
                charIndex--;
            } else {
                charIndex++;
            }
            setTimeout(typeEffect, isDeleting ? 50 : 100);
        }
    }

    // Start the animation
    typeEffect();
}

// Quiz Score Handler
document.body.addEventListener('htmx:afterRequest', function(evt) {
    const triggeredEvents = evt.detail.headers?.['HX-Trigger'];
    if (!triggeredEvents) return;
    
    try {
        const events = JSON.parse(triggeredEvents);
        if (events.updateScore) {
            const score = document.getElementById('currentScore');
            const attempted = document.getElementById('questionsAttempted');
            const correct = document.getElementById('correctAnswers');
            
            if (!score || !attempted || !correct) {
                console.error('Could not find required quiz elements');
                return;
            }
            
            attempted.textContent = parseInt(attempted.textContent) + 1;
            if (events.updateScore.correct) {
                correct.textContent = parseInt(correct.textContent) + 1;
                score.textContent = parseInt(score.textContent) + 10;
            }
        }
    } catch (error) {
        console.error('Error handling quiz score update:', error);
    }
});