export default class TodoItem {
    constructor(text, id = null) {
        this.id = id || Date.now();
        this.text = text;
    }

    render() {
        const li = document.createElement('li');
        li.className = 'todo-item';
        li.innerHTML = `
            <span>${this.text}</span>
            <button>Xóa</button>
        `;
        return li;
    }
}
