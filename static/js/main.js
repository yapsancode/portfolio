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