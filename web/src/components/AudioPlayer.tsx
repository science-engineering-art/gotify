
export function AudioPlayer(props: { src: string}) {
  return (
    <audio id="audio" controls={true} autoPlay={true}>
      <source src={props.src} type="audio/mpeg" />
    </audio>
  );
}
