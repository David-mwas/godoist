// STARTER CODE:

import { Button, Flex, Input, Spinner } from "@chakra-ui/react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useState } from "react";
import { IoMdAdd } from "react-icons/io";
import { BASE_URL } from "../App";

const TodoForm = () => {
  const queryClient = useQueryClient();

  const [newTodo, setNewTodo] = useState("");

  const { mutate: createTodo, isPending: isCreatingTodo } = useMutation({
    mutationKey: ["createTodo"],
    mutationFn: async (e: React.FormEvent) => {
      try {
        e.preventDefault();

        const res = await fetch(`${BASE_URL}/todo`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            body: newTodo,
          }),
        });
        const data = await res.json();

        if (!res.ok) {
          throw new Error(data?.error || "something went wrong");
        }

        return data;
      } catch (error) {
        console.error(error);
      } finally {
        setNewTodo("");
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["todos"] });
    },

    onError: (error) => {
      alert(error?.message);
    },
  });
  return (
    <form onSubmit={createTodo}>
      <Flex gap={2}>
        <Input
          type="text"
          value={newTodo}
          onChange={(e) => setNewTodo(e.target.value)}
          ref={(input) => input && input.focus()}
        />
        <Button
          mx={2}
          type="submit"
          _active={{
            transform: "scale(.97)",
          }}
        >
          {isCreatingTodo ? <Spinner size={"xs"} /> : <IoMdAdd size={30} />}
        </Button>
      </Flex>
    </form>
  );
};
export default TodoForm;
