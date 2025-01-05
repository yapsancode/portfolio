// internal/handlers/experience.go
package handlers

import (
	"fmt"
	"net/http"
)

func ExperienceHandler(w http.ResponseWriter, r *http.Request) {
	html := `
    <section>
        <h2 class="text-2xl font-bold mb-4">Working Experience</h2>
        <div class="space-y-4">
            <div class="bg-white p-4 rounded shadow">
                <h3 class="font-semibold">Software Developer at ABC Co.</h3>
                <p>Developed scalable web applications and collaborated with cross-functional teams.</p>
            </div>
            <div class="bg-white p-4 rounded shadow">
                <h3 class="font-semibold">Intern at XYZ Inc.</h3>
                <p>Worked on IoT projects and implemented innovative solutions for client needs.</p>
            </div>
        </div>
    </section>`

	fmt.Fprint(w, html)
}
