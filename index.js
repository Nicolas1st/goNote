const onlyUseLocalstorage = true; 
const isBackendAvailable = async function() {
	try {
		await fetch(`/notes/`);
		onlyUseLocalstorage = false;
	} catch (error) {
	}
}
isBackendAvailable();


// creating an HTML component for note
const components = {
    newNoteComponent: function (noteTitle, noteContent, noteID) {
        // note itself
        const note = document.createElement("div");
        note.classList.add("note");

        // note's title
        const title = document.createElement("h2");
        title.innerText = noteTitle;
        title.classList.add("title");
        title.setAttribute("contenteditable", "true");

        // note's content
        const content = document.createElement("p");
        content.innerText = noteContent;
        content.setAttribute("contenteditable", "true");
        content.classList.add("content");

        // note's remove button
        const removeButton = document.createElement("div");
        removeButton.classList.add("remove-button");

        // note's save changes button
        const saveChangesButton = document.createElement("div");
        saveChangesButton.classList.add("save-changes-button");
        saveChangesButton.innerText = "Save";

        saveChangesButton.addEventListener("click", () => {
            note.removeChild(saveChangesButton);
            note.dataset.modified = false;

            note.dispatchEvent(
                new Event("notechange", {
                    bubbles: true,
                })
            );
        });

        // assembling
        note.appendChild(title);
        note.appendChild(content);
        note.appendChild(removeButton);

        // setting data attributes
        note.dataset.id = noteID;
        note.dataset.modified = false;

        // handling contents change event
        note.addEventListener("input", (e) => {
            if (!note.dataset.modified) {
                return;
            }

            note.dataset.modified = true;
            note.appendChild(saveChangesButton);
        });

        return note;
    },
};

// the presentation object presents changes (renders) what's requested
const presentation = {
    displayNoteComponents: function (replace, ...notes) {
        const notesContainer = document.querySelector(".notes-container");

        if (replace) {
            notesContainer.innerHTML = "";
        }

        // adding notes
        notes.forEach((note) => {
            notesContainer.appendChild(
                components.newNoteComponent(note.Title, note.Content, note.NoteID)
            );
        });
    },

    removeNoteComponent: function (note) {
        const notesContainer = document.querySelector(".notes-container");

        notesContainer.removeChild(note);
    },

    updateNoteComponent: function (note) {
        const notesContainer = document.querySelector(".notes-container");

        const newNote = components.newNoteComponent(note.Title, note.Content, note.NoteID);
        const oldNote = notesContainer.querySelector(`[data-id="${note.NoteID}"]`);

        notesContainer.replaceChild(newNote, oldNote);
    },
};

/* all fucntions defined in the network object work in the following way */
/* they make requests to the api */
/* if succesfull they return a json */
/* otherwise they error out */
/* it's the calling code responsibility to check for errors */
const network = {
    loadNotes: async function() {
        const response = await fetch(`/notes/`);

        return JSON.parse(JSON.stringify(await response.json()));
    },

    createNote: async function (note) {
        const response = await fetch(`/notes/`, {
            method: "POST",
            header: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
				NoteID: note.NoteID,
                Title: note.Title,
                Content: note.Content
            }),
        });

        return response.json();
    },

    deleteNote: async function (noteID) {
        const response = await fetch(`/notes/${noteID}/`, {
            method: "DELETE",
        });

        return await response.json();
    },

    updateNote: async function (noteID, noteTitle, noteContent) {
        const response = await fetch(`/notes/${noteID}/`, {
            method: "PUT",
            body: JSON.stringify({
                Title: noteTitle,
                Content: noteContent,
            }),
        });

        return await response.json();
    },
}

const main = function(components, presentation, network) {
	// loading available notes from the api on load
	document.addEventListener("DOMContentLoaded", async () => {
		let notes = [];

		if (onlyUseLocalstorage) {
			notesFromStorage = {...localStorage};
			for (let NoteID in notesFromStorage) {
				notes.push(JSON.parse(notesFromStorage[NoteID]))
			}
		} else {
			try {
				notes = await network.loadNotes();
				notes.forEach((note) => {
					localStorage.setItem(note.NoteID, JSON.stringify(note))
				});
			} catch (error) {
			}
		}

		presentation.displayNoteComponents(false, ...notes);
	});

	const newNoteForm = document.querySelector(".form");
	newNoteForm.addEventListener("submit", async (e) => {
		e.preventDefault();
		if (!newNoteForm.reportValidity()) {
			return;
		}

		const titleElement = newNoteForm.querySelector("#title");
		const contentElement = newNoteForm.querySelector("#content");

		let createdNote = {
			// using a timestamp as an id
			// because the author is being used as namespace
			// and also the program is designed for human use
			// it's very unlikey for a human to create several notes at an exact same millisecond
			NoteID: Date.now().toString(), 
			Title: titleElement.value,
			Content: contentElement.value,
		};

		if (!onlyUseLocalstorage) {
			try {
				await network.createNote(createdNote);
			} catch (error) {
			}
		}

		localStorage.setItem(createdNote.NoteID, JSON.stringify(createdNote))

		// clearing the form fields
		titleElement.value = "";
		contentElement.value = "";

		// displaying the note
		presentation.displayNoteComponents(false, createdNote);
	});

	// handling Enter key presses on form
	newNoteForm.addEventListener("keypress", (e) => {
		if (e.key === "Enter") {
			// fire submit form event if shift and enter are pressed together
			if (e.shiftKey) {
				e.target.form.dispatchEvent(
					new Event("submit", { cancelable: true })
				);
				e.preventDefault();
			}

			// prevent the input element from firing the submit event
			if (e.target instanceof HTMLInputElement) {
				e.preventDefault();

				const nextEl = e.target.parentElement.nextSibling.nextSibling.querySelector(".form-input");
				if (nextEl instanceof HTMLElement) {
					nextEl.focus();
				}
			}
		}
	});

	// removing a note
	const notesContainer = document.querySelector(".notes-container");
	notesContainer.addEventListener("click", async (e) => { const removeButton = e.target;
		// checking whether it is a remove button
		if (!removeButton.classList.contains("remove-button")) {
			return;
		}

		// find the note that that has this button
		const note = removeButton.closest(".note");

		if (!onlyUseLocalstorage) {
			try {
				const removedNote = await network.deleteNote(note.dataset.id);
			} catch (error) {
			}
		}

		localStorage.removeItem(note.dataset.id)

		// removing the note from the dom if the request was successful
		presentation.removeNoteComponent(note);
	});

	document.addEventListener("notechange", async (e) => {
		const note = e.target;

		const updatedNote = {
			NoteID: note.dataset.id,
			Title: note.querySelector(".title").innerText,
			Content: note.querySelector(".content").innerText
		};

		if (!onlyUseLocalstorage) {
			try {
				await network.updateNote(
					updatedNote.NoteID,
					updatedNote.Title,
					updatedNote.Content
				);
			} catch (error) {
			}
		}

		localStorage.setItem(updatedNote.NoteID, JSON.stringify(updatedNote))

		presentation.updateNoteComponent(updatedNote);
	});
}

main(components, presentation, network)
