
export function AudioPlayer(props: { src: string}) {
  return (
    <audio autoPlay={true} controls>
      <source src={props.src} type="audio/mpeg" />
    </audio>
  );
}
