import { displayNoteComponents, removeNoteComponent, updateNoteComponent } from "./presentation.js";
import { loadNotes, createNote, deleteNote, updateNote } from "./network.js";

// loading available notes from the api on load
document.addEventListener("DOMContentLoaded", async () => {
    try {
        const notes = await loadNotes("Nicolas");
        displayNoteComponents(false, ...notes);
    } catch (error) {
        console.log("Could not connect to the api");
        console.log(error);
    }
});

const newNoteForm = document.querySelector(".form");
newNoteForm.addEventListener("submit", async (e) => {
    e.preventDefault();
    if (!newNoteForm.reportValidity()) {
        return;
    }

    const titleElement = newNoteForm.querySelector("#title");
    const contentElement = newNoteForm.querySelector("#content");

    // making a request to create a note
    // if successful displaying changes on the page
    let createdNote;
    try {
        createdNote = await createNote(
            "Nicolas",
            titleElement.value,
            contentElement.value
        );
    } catch (error) {
        console.log("Error occured when making a request to the api");
        console.log(error);
    }

    // clearing the form fields
    titleElement.value = "";
    contentElement.value = "";

    // displaying the note
    displayNoteComponents(false, createdNote);
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

    // making the api request
    try {
        const removedNote = await deleteNote(note.dataset.id);
    } catch (error) {
        console.log("Could not make the request to delete the note");
        console.log(error);
        return;
    }

    // removing the note from the dom if the request was successful
    removeNoteComponent(note);
});

document.addEventListener("notechange", async (e) => {
    const note = e.target;

    const id = note.dataset.id;
    const title = note.querySelector(".title").innerText;
    const content = note.querySelector(".content").innerText;

    let updatedNote;
    try {
        updatedNote = await updateNote(id, title, content);
    } catch (error) {
        console.log("Could not make the request to delete the note");
        console.log(error);
        return;
    }

    updateNoteComponent(updatedNote);
});