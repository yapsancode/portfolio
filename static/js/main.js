// static/js/main.js

// Audio Player Script
function initializeAudioPlayer() {
    const playButton = document.getElementById('playAudio');
    const audioIcon = document.getElementById('audioIcon');
    const audioTooltip = document.getElementById('audioTooltip');
    const audio = document.getElementById('welcomeAudio');

    if (playButton && audioIcon && audioTooltip && audio) {
        // Remove any existing listeners to prevent duplicates
        playButton.removeEventListener('click', handleAudioClick);

        playButton.addEventListener('click', handleAudioClick);

        audio.addEventListener('ended', () => {
            audioIcon.classList.remove('fa-pause');
            audioIcon.classList.add('fa-play');
            audioTooltip.textContent = 'Play Welcome Audio';
            playButton.classList.remove('scale-110');
        });
    }
}

function handleAudioClick() {
    const playButton = document.getElementById('playAudio');
    const audioIcon = document.getElementById('audioIcon');
    const audioTooltip = document.getElementById('audioTooltip');
    const audio = document.getElementById('welcomeAudio');

    if (audio.paused) {
        audio.play().catch(error => console.error('Error playing audio:', error));
        audioIcon.classList.remove('fa-play');
        audioIcon.classList.add('fa-pause');
        audioTooltip.textContent = 'Pause Welcome Audio';
    } else {
        audio.pause();
        audioIcon.classList.remove('fa-pause');
        audioIcon.classList.add('fa-play');
        audioTooltip.textContent = 'Play Welcome Audio';
    }

    playButton.classList.add('scale-110');
    setTimeout(() => playButton.classList.remove('scale-110'), 200);
}

// Tech Stack Cards Functionality
const techDetails = {
    'go': `
        <div class="bg-slate-700/30 p-4 rounded-lg">
            <h2 class="text-xl font-semibold mb-3 flex items-center">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 mr-2 text-blue-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
                </svg>
                Why Go?
            </h2>
            <ul class="list-none space-y-2 text-gray-300">
                <li class="flex items-start"><span class="text-green-400 mr-2">•</span>Blazing fast—compiled, lean, and built for concurrency. High-level ease, low-level punch.</li>
                <li class="flex items-start"><span class="text-green-400 mr-2">•</span>Future-proof. Microservices and cloud-native trends eat this up—think Kubernetes vibes.</li>
            </ul>
        </div>
    `,
    'htmx': `
        <div class="bg-slate-700/30 p-4 rounded-lg">
            <h2 class="text-xl font-semibold mb-3 flex items-center">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 mr-2 text-blue-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4" />
                </svg>
                Why HTMX?
            </h2>
            <ul class="list-none space-y-2 text-gray-300">
                <li class="flex items-start"><span class="text-green-400 mr-2">•</span>Lightweight magic—swaps heavy JS frameworks for HTML-driven interactivity. Less bloat, more speed.</li>
                <li class="flex items-start"><span class="text-green-400 mr-2">•</span>Server-side simplicity. Pairs with Go like a dream—no client-side overload.</li>
            </ul>
        </div>
    `,
    'vanilla-js': `
        <div class="bg-slate-700/30 p-4 rounded-lg">
            <h2 class="text-xl font-semibold mb-3 flex items-center">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 mr-2 text-yellow-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
                </svg>
                Why Vanilla JS?
            </h2>
            <ul class="list-none space-y-2 text-gray-300">
                <li class="flex items-start"><span class="text-green-400 mr-2">•</span>No framework fat—pure, lean JS for total control and zero bloat.</li>
                <li class="flex items-start"><span class="text-green-400 mr-2">•</span>Plays nice with HTMX. Handles the small, dynamic bits without overcomplicating the stack.</li>
            </ul>
        </div>
    `,
    'tailwind': `
        <div class="bg-slate-700/30 p-4 rounded-lg">
            <h2 class="text-xl font-semibold mb-3 flex items-center">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 mr-2 text-teal-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zm0 0h12a2 2 0 002-2v-4a2 2 0 00-2-2h-2.343M11 7.343l1.657-1.657a2 2 0 012.828 0l2.829 2.829a2 2 0 010 2.828l-8.486 8.485M7 17h.01" />
                </svg>
                Why TailwindCSS?
            </h2>
            <ul class="list-none space-y-2 text-gray-300">
                <li class="flex items-start"><span class="text-green-400 mr-2">•</span>Custom vibes—ditches Bootstrap's generic look for pixel-perfect control.</li>
                <li class="flex items-start"><span class="text-green-400 mr-2">•</span>Rapid styling. Inline utility classes mean no CSS ping-pong—build fast, tweak faster.</li>
            </ul>
        </div>
    `,
    'mysql': `
        <div class="bg-slate-700/30 p-4 rounded-lg">
            <h2 class="text-xl font-semibold mb-3 flex items-center">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 mr-2 text-yellow-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7v10c0 2 1 3 3 3h10c2 0 3-1 3-3V7c0-2-1-3-3-3H7c-2 0-3 1-3 3z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 11h6m-3-3v6" />
                </svg>
                Why MySQL?
            </h2>
            <ul class="list-none space-y-2 text-gray-300">
                <li class="flex items-start"><span class="text-green-400 mr-2">•</span>Beginner-friendly—simple to set up, tons of docs, and battle-tested for small gigs.</li>
                <li class="flex items-start"><span class="text-green-400 mr-2">•</span>Reliable relational backbone. Perfect for structured data without overcomplicating things.</li>
            </ul>
        </div>
    `
};

function initializeTechCards() {
    const tags = document.querySelectorAll('#stack-tags span');
    const container = document.getElementById('tech-card-container');
    
    if (tags.length > 0 && container) {
        // Remove any existing listeners to prevent duplicates
        tags.forEach(tag => {
            tag.removeEventListener('click', handleTagClick);
            tag.addEventListener('click', handleTagClick);
        });
        
        // Show the first tech card by default
        if (tags[0]) {
            const defaultTech = tags[0].getAttribute('data-tech');
            if (defaultTech && techDetails[defaultTech]) {
                container.innerHTML = techDetails[defaultTech];
                container.classList.remove('opacity-0');
                container.classList.add('opacity-100');
            }
        }
    }
}

function handleTagClick(event) {
    const tech = event.currentTarget.getAttribute('data-tech');
    const container = document.getElementById('tech-card-container');
    
    if (tech && techDetails[tech] && container) {
        // Add a fade-out effect
        container.classList.add('opacity-0');
        
        // Change content after a short delay for the fade effect
        setTimeout(() => {
            container.innerHTML = techDetails[tech];
            container.classList.remove('opacity-0');
            container.classList.add('opacity-100');
        }, 150);
    }
}

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

// Sidebar Toggle Script
function initializeSidebar() {
    const sidebar = document.getElementById('sidebar');
    const toggleSidebar = document.getElementById('toggleSidebar');

    if (sidebar && toggleSidebar) {
        toggleSidebar.removeEventListener('click', handleSidebarToggle);
        toggleSidebar.addEventListener('click', handleSidebarToggle);
    }
}

function handleSidebarToggle() {
    const sidebar = document.getElementById('sidebar');
    console.log('Sidebar toggle clicked');
    sidebar.classList.toggle('w-16');
    sidebar.classList.toggle('w-64');
}

// Initialize all features
function initializeAllFeatures() {
    initializeAudioPlayer();
    initializeTechCards();
    startTextAnimation();
    initializeSidebar();
}

// Initialize on both events
document.addEventListener('DOMContentLoaded', initializeAllFeatures);
document.addEventListener('htmx:load', initializeAllFeatures);
document.body.addEventListener('htmx:afterSwap', function (evt) {
    if (evt.detail.target.id === 'mainContent') {
        initializeAllFeatures();
    }
});