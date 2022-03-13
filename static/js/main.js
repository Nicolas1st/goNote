// creating components

function NewNoteComponent(noteTitle, noteContent) {
    // note itself
    const note = document.createElement("div");
    note.classList.add("note");

    // note's title
    const title = document.createElement("h2");
    title.innerText = noteTitle;
    title.classList.add("title");

    // note's content
    const content = document.createElement("p");
    content.innerText = noteContent;
    content.classList.add("content");

    // assembling
    note.appendChild(title);
    note.appendChild(content);

    return note;
}

// main
const notes = document.querySelector(".notes-container");

const noteTitle= "Super Cool Note";
const noteContent = "Engaging Contents";

Array(10).fill().map(() => {
    const note = NewNoteComponent(noteTitle, noteContent);
    notes.appendChild(note);
});