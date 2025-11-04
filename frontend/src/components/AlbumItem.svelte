<script lang="ts">
  import type { Album } from "../types/album";

  export let album: Album;
  export let handleRemove: (id: number) => void;
</script>

<li>
  <strong>{album.title} by {album.artist} - ${album.price}</strong>
  <button on:click={() => handleRemove(album.id)}>Remove</button>
</li>
