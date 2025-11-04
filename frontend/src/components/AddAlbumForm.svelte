<script lang="ts">
  import type { Album } from "../types/album";

  export let albums: Album[];

  let formData: Album = {
    id: 0,
    title: "",
    artist: "",
    price: 0,
  };

  async function onSubmit() {
    const res = await fetch(`http://localhost:8080/albums/add`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(formData),
    });

    const newAlbumID: number = await res.json();

    if (res.ok) {
      albums = [...albums, { ...formData, id: newAlbumID }];
    }

    formData = {
      id: 0,
      title: "",
      artist: "",
      price: 0,
    };
  }
</script>

<form
  class="flex flex-col gap-2 m-20"
  on:submit|preventDefault={onSubmit}
>
  <div
    class="flex flex-col gap-2 w-100 *:flex *:gap-2 *:justify-between *:*:bg-gray-400 *:*:pl-2"
  >
    <label
      >Title:
      <input
        placeholder="title here"
        type="text"
        bind:value={formData.title}
        required
      />
    </label>
    <label>
      Artist:
      <input
        placeholder="artist here"
        type="text"
        bind:value={formData.artist}
        required
      />
    </label>
    <label>
      Price:
      <input
        type="number"
        bind:value={formData.price}
        min="0"
        step="0.01"
        required
      />
    </label>
  </div>
  <button type="submit" class="bg-black text-white w-1/4 h-10"
    >Submit</button
  >
</form>
