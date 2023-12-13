<script setup>
// Эта часть выполняется при первоначальной загрузке компонента
</script>

<style>
  .note {
    cursor: pointer;
    padding: 1rem 0.25rem;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  .note:hover {
    background-color: beige;
  }
  .note.selected {
    background-color: #19a9cd;
  }
  .task {
    padding: 1rem 0.25rem;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
</style>

<template>
  <div class="row">
    <div class="col-md-6">
      <div v-for="note in notes" @click="current_note = note" :class="{ 'note': true, 'selected':current_note == note }">
        <div v-if="!note.isEditMode">{{ note.name }}</div>
        <div v-if="note.isEditMode">
          <input type="text" class="form-control" v-model="note.name" @keydown.enter="note.isEditMode = false">
        </div>
        <div>
          <button v-if="note.isEditMode" class="btn btn-sm btn-success" @click="note.isEditMode = false"><i class="bi bi-check"></i></button>
          <button v-if="!note.isEditMode" class="btn btn-sm btn-primary" @click="note.isEditMode = true"><i class="bi bi-pencil"></i></button>
          <button v-if="!note.isEditMode" class="btn btn-sm btn-danger" @click="removeNote(task)"><i class="bi bi-trash"></i></button>
        </div>
      </div>
      <button class="btn-primary w-100" @click="addNote">Добавить заметку</button>
    </div>
    <div class="col-md-6">
      <template v-if="current_note != null">
        <div v-for="task in current_note.tasks" class="task">
          <div v-if="!task.isEditMode">
            <input type="checkbox" v-model="task.checked">
            {{ task.name }}
          </div>
          <div v-if="task.isEditMode">
            <input type="text" class="form-control" v-model="task.name" @keydown.enter="task.isEditMode = false">
          </div>
          <div>
            <button v-if="task.isEditMode" class="btn btn-sm btn-success" @click="task.isEditMode = false"><i class="bi bi-check"></i></button>
            <button v-if="!task.isEditMode" class="btn btn-sm btn-primary" @click="task.isEditMode = true"><i class="bi bi-pencil"></i></button>
            <button v-if="!task.isEditMode" class="btn btn-sm btn-danger" @click="removeTask(task)"><i class="bi bi-trash"></i></button>
          </div>
        </div>
        <button class="btn-primary w-100" @click="addTask">Добавить задачу</button>
      </template>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      current_note: null,
      notes: [],
    }
  },
  computed: {
  },
  methods: {
    removeTask(task) {
      if (!confirm("Вы уверены, что хотите удалить задачу?")) return;

      const index = this.current_note.tasks.indexOf(task);
      this.current_note.tasks.splice(index, 1);
    },
    removeNote(note) {
      if (!confirm("Вы уверены, что хотите удалить замтеку?")) return;

      const index = this.notes.indexOf(note);
      this.notes.splice(index, 1);
    },
    addNote() {
      this.notes.push({
        name: 'Новая заметка',
        tasks: [],
        isEditMode: false,
      });
    },
    addTask() {
      this.current_note.push({
        name: 'Новая задача',
        checked: false,
        isEditMode: false,
      });
    }
  },
}
</script>
