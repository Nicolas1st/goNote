export function newNoteComponent(noteTitle, noteContent, noteID) {
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

    content.appendChild(newSaveButton());

    // note's remove button
    const removeButton = document.createElement("div");
    removeButton.classList.add("remove-button");

    // assembling
    note.appendChild(title);
    note.appendChild(content);
    note.appendChild(removeButton);

    // setting data attributes
    note.dataset.id = noteID; 

    return note;
}

export function newSaveButton() {
    const button = document.createElement("div");
    button.classList.add("save-changes-button");
    button.innerText = "Save";

    return button;
}