
export function AudioPlayer(props: { src: string}) {
  console.log(props.src)
  return (
    <audio autoPlay={true} controls>
      <source src={props.src} type="audio/mpeg" />
    </audio>
  );
}
