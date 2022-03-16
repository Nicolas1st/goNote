export const components = {
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
