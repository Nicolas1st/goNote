import { displayNoteComponents, removeNoteComponent } from "./presentation.js";
import { loadNotes, createNote, deleteNote } from "./network.js";

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
        const removedNote = await deleteNote("Nicolas", note.dataset.id);
    } catch (error) {
        console.log("Could not make the request to delete the note");
        console.log(error);
        return;
    }

    // removing the note from the dom if the request was successful
    removeNoteComponent(note);
});
