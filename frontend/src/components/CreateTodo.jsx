import { useState } from "react";

export function CreateTodo() {
    const [title, setTitle] = useState("");
    const [description, setDescription] = useState("");

    return (
        <div>
            <input
                type="text"
                placeholder="title"
                onChange={(e) => setTitle(e.target.value)}
            /><br />
            <input
                type="text"
                placeholder="description"
                onChange={(e) => setDescription(e.target.value)}
            /><br />
            <button onClick={() => {
                fetch("http://127.0.0.1:8080/todo", {
                    method: "POST",
                    body: JSON.stringify({
                        title: title,
                        description: description
                    }),
                    headers: {
                        "Content-Type": "application/json"
                    }
                })
                    .then(async function (res) {
                        const json = await res.json();
                        alert("todo added");
                    })
                    .catch(function (error) {
                        console.error("Error adding todo:", error);
                        alert("Failed to add todo");
                    });
            }}>Add a todo</button>
        </div>
    );
}
