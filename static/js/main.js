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

// Text Animation
const roles = [
    "Software Developer",
    'UI/UX Enthusiast',
    'Problem Solver',
    "Software Engineer",
    "Full-Stack Developer",
    "Mobile App Developer",
];

function startTextAnimation() {
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
}

// Initialize the animation when the page loads
startTextAnimation();

// Reinitialize the animation when HTMX swaps content
document.body.addEventListener('htmx:afterSwap', function (evt) {
    if (evt.detail.target.id === 'mainContent') {
        startTextAnimation();
    }
});

// Sidebar Toggle Script
const sidebar = document.getElementById('sidebar');
const toggleSidebar = document.getElementById('toggleSidebar');

toggleSidebar?.addEventListener('click', () => {
    console.log('Sidebar toggle clicked');
    sidebar.classList.toggle('w-16');
    sidebar.classList.toggle('w-64');
});

// Quiz Dialog Handler
document.body.addEventListener('showUsernameDialog', function () {
    const dialog = document.getElementById('usernameDialog');
    if (dialog) {
        dialog.showModal();
    }
});

document.body.addEventListener('hideUsernameDialog', function () {
    const dialog = document.getElementById('usernameDialog');
    if (dialog) {
        dialog.close();
    }
});

// Quiz Score Handler
document.body.addEventListener('htmx:afterOnLoad', function (evt) {
    if (evt.detail.elt.matches('[hx-post="/quiz/answer"]')) {
        try {
            const data = JSON.parse(evt.detail.xhr.response);
            if (data.updateScore) {
                const score = document.getElementById('currentScore');
                const attempted = document.getElementById('questionsAttempted');
                const correct = document.getElementById('correctAnswers');

                if (data.updateScore.correct) {
                    score.textContent = (parseInt(score.textContent) || 0) + 1;
                    correct.textContent = score.textContent;
                }
                attempted.textContent = (parseInt(attempted.textContent) || 0) + 1;
            }
        } catch (e) {
            console.error('Error parsing response:', e);
        }
    }
});

document.body.addEventListener('quizCompleteDialog', () => {
    const content = document.getElementById('quizContent');
    if (content) {
        content.innerHTML = `
            <div class="p-6 rounded-lg shadow-md bg-green-100 text-green-900">
                <h2 class="text-2xl font-bold mb-4">Congrats!</h2>
                <p>You have completed the quiz!</p>
            </div>`;
    }
});

// static/js/main.js
document.body.addEventListener('updateScore', function () {
    const score = document.getElementById('currentScore');
    if (score) {
        score.textContent = parseInt(score.textContent || 0) + 1;
    }
});

document.body.addEventListener('updateAttempts', function () {
    const attempts = document.getElementById('questionsAttempted');
    if (attempts) {
        attempts.textContent = parseInt(attempts.textContent || 0) + 1;
    }
});