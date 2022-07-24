<script>
    import { onMount } from "svelte";
    import Pusher from "pusher-js";

    let username = "username";
    let message = "";
    let messages = [];

    onMount(() => {
        Pusher.logToConsole = true;

        const pusher = new Pusher("659c70e291e6e4b7fb80", {
            cluster: "ap1",
        });

        const channel = pusher.subscribe("chat");
        channel.bind("message", (data) => {
            messages = [...messages, data];
        });
    });

    async function submit() {
        await self.fetch("http://localhost:8000/api/messages", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
                username,
                message,
            }),
        });
        message = "";
    }
</script>

<div class="container">
    <div class="d-flex flex-column align-items-stretch flex-shrink-0 bg-white">
        <div
            class="d-flex align-items-center flex-shrink-0 p-3 link-dark text-decoration-none border-bottom"
        >
            <input class="fs-5 fw-semibold" type="text" bind:value={username} />
        </div>
        <div class="list-group list-group-flush border-bottom scrollarea">
            {#each messages as msg}
                <div class="list-group-item list-group-item-action py-3 lh-sm">
                    <div
                        class="d-flex w-100 align-items-center justify-content-between"
                    >
                        <strong class="mb-1">{msg.username}</strong>
                    </div>
                    <div class="col-10 mb-1 small">{msg.message}</div>
                </div>
            {/each}
        </div>
    </div>

    <form on:submit|preventDefault={submit}>
        <input
            class="form-control"
            placeholder="Write a message"
            bind:value={message}
        />
        <input type="submit" />
    </form>
</div>

<style>
    .scrollarea {
        min-height: 500px;
    }
</style>
