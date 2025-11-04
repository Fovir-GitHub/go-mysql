<script lang="ts">
  import { onMount } from "svelte";
  import AlbumList from "./components/AlbumList.svelte";
  import type { Album } from "./types/album";
  import AddAlbumForm from "./components/AddAlbumForm.svelte";

  let albums: Album[] = [];
  let error: string | null = null;

  onMount(async () => {
    try {
      const res = await fetch("http://localhost:8080/albums/all");
      if (!res.ok) {
        throw new Error(`HTTP ${res.status}`);
      }
      albums = (await res.json()) as Album[];
    } catch (e: any) {
      error = e.message;
    }
  });
</script>

<main>
  <AddAlbumForm bind:albums />

  {#if error}
    <p style="color: red">{error}</p>
  {:else if albums.length === 0}
    <p>No album</p>
  {:else}
    <AlbumList bind:albums />
  {/if}
</main>
