let lastID = 2;

export function rootReducer(state = { isLoading: false, tasks: [] }, action) {
  switch (action.type) {
    case "SET_TASKS":
      return { ...state, tasks: action.tasks };
    case "ADD_TASK":
      return {
        ...state,
        tasks: [...state.tasks, { id: lastID++, title: action.title }]
      };
    case "LOADING":
      return { ...state, isLoading: true };
    case "LOADED":
      return { ...state, isLoading: false };
  }
  return state;
}
