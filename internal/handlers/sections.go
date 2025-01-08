// internal/handlers/sections.go
package handlers

import (
	"fmt"
	"net/http"
)

func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	html := `
    <section>
        <h2 class="text-2xl font-bold mb-4">Projects</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div class="bg-white p-4 rounded shadow">
                <h3 class="font-semibold">Project 1</h3>
                <p>A web app for task management with real-time collaboration.</p>
            </div>
            <div class="bg-white p-4 rounded shadow">
                <h3 class="font-semibold">Project 2</h3>
                <p>An IoT-based system for smart home automation.</p>
            </div>
        </div>
    </section>`
	fmt.Fprint(w, html)
}

func ContactSectionHandler(w http.ResponseWriter, r *http.Request) {
	html := `
    <section>
        <h2 class="text-2xl font-bold mb-4">Contact Me</h2>
        <form method="POST" action="/contact-submit" class="space-y-4 bg-white p-4 rounded shadow">
            <div>
                <label for="name" class="block font-semibold">Name:</label>
                <input type="text" id="name" name="name" required class="w-full p-2 border rounded">
            </div>
            <div>
                <label for="email" class="block font-semibold">Email:</label>
                <input type="email" id="email" name="email" required class="w-full p-2 border rounded">
            </div>
            <div>
                <label for="message" class="block font-semibold">Message:</label>
                <textarea id="message" name="message" required class="w-full p-2 border rounded h-32"></textarea>
            </div>
            <button type="submit" class="bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600">Send</button>
        </form>
    </section>`
	fmt.Fprint(w, html)
}

func AMASectionHandler(w http.ResponseWriter, r *http.Request) {
	html := `
    <section>
        <h2 class="text-2xl font-bold mb-4">Ask Me Anything</h2>
        <form method="POST" action="/ama-submit" class="space-y-4 bg-white p-4 rounded shadow">
            <div>
                <label for="question" class="block font-semibold">Your Anonymous Question:</label>
                <textarea id="question" name="question" required class="w-full p-2 border rounded h-32"></textarea>
            </div>
            <button type="submit" class="bg-yellow-500 text-white px-4 py-2 rounded hover:bg-yellow-600">Submit Anonymously</button>
        </form>
    </section>`
	fmt.Fprint(w, html)
}
