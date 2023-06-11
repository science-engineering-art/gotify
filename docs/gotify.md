# Informe del Proyecto de Sistemas Distribuidos: Gotify

### Integrantes:
- Rolando Sánchez Ramos C411
- Javier Villar Alonso C412
- Leandro Rodriguez Llosa C411

## Introducción:

El continuo avance de la tecnología ha cambiado radicalmente la forma en que consumimos contenido multimedia, especialmente la música. En la actualidad, cada vez más personas recurren a plataformas de streaming para satisfacer sus necesidades de entretenimiento musical. Sin embargo, el crecimiento y desarrollo de estos servicios no debe detenerse aquí, y por ello presentamos nuestro proyecto, Gotify.

Gotify es una propuesta innovadora para un sistema de streaming de música basado en tecnología de sistemas distribuidos. La idea central de este proyecto es desarrollar una plataforma que permita a cualquier host conectarse y reproducir música, brindando al mismo tiempo un robusto sistema de almacenamiento y distribución de música.

El sistema Gotify ofrecerá una interfaz de usuario sencilla e intuitiva, que permitirá a los usuarios subir, listar, buscar y reproducir música. Cada pista musical contará con metadatos asociados, como género, autor, álbum y nombre, para que los usuarios puedan realizar búsquedas orientadas según sus preferencias.

El núcleo de Gotify está formado por un conjunto de nodos, cada uno con una función específica, que trabajan juntos para garantizar la disponibilidad y la resistencia del sistema. Este diseño distribuido permitirá que el sistema siga funcionando siempre que haya al menos un nodo disponible por cada rol, y preservará los datos incluso si un nodo de almacenamiento falla.

Una característica fundamental del sistema Gotify es su tolerancia a fallos. Incluso en el caso de una partición del sistema, donde se divida en dos o más subsistemas debido a la pérdida de nodos o enlaces entre nodos, el sistema será capaz de reconectarse y funcionar como una única entidad una vez que se restablezca la comunicación entre los nodos.

Finalmente, el sistema Gotify tendrá la restricción de que, una vez que comienza la reproducción de una canción, no debe interrumpirse debido a un fallo del sistema. Con este enfoque, el proyecto Gotify aspira a ofrecer una experiencia de escucha ininterrumpida y de alta calidad a todos sus usuarios.

En resumen, el proyecto Gotify representa un avance significativo en el campo del streaming de música, proporcionando un sistema altamente disponible y resistente a fallos, que ofrece una experiencia de usuario excelente.

### Objetivos:
1. **Desarrollar una interfaz de usuario intuitiva y accesible:** Crear una plataforma que permita a los usuarios subir, listar, buscar y reproducir música de manera sencilla y rápida, proporcionando una experiencia de usuario atractiva y funcional.

2. **Implementar un sistema basado en metadatos:** Cada pista de música tendrá metadatos asociados como género, autor, álbum y nombre. Este enfoque permitirá búsquedas eficientes y personalizadas, mejorando la experiencia del usuario.

3. **Crear un sistema distribuido resiliente y tolerante a fallos:** Diseñar un conjunto de nodos, cada uno con una función específica, que trabajen de manera conjunta para garantizar la disponibilidad y resistencia del sistema incluso en caso de fallos o pérdida de nodos.

4. **Garantizar la disponibilidad permanente del servicio:** El sistema estará diseñado para seguir funcionando siempre que haya al menos un nodo disponible por cada rol, lo que permitirá la continuidad del servicio bajo cualquier circunstancia.

5. **Implementar un sistema de almacenamiento seguro y robusto:** El sistema no deberá perder datos, incluso en caso de fallo de un nodo de almacenamiento. Este enfoque asegurará la integridad y la seguridad de la música y los metadatos asociados.

6. **Desarrollar un sistema capaz de recuperarse de particiones:** En caso de una partición del sistema, este deberá ser capaz de reconectarse y funcionar como un solo sistema una vez que se restablezca la comunicación entre los nodos.

7. **Ofrecer una reproducción de música ininterrumpida:** Independientemente de las condiciones del sistema, una vez que comienza la reproducción de una canción, no deberá interrumpirse. Esto proporcionará una experiencia de escucha de alta calidad a los usuarios.

## Desarrollo:

### Descripción General:

El sistema Gotify se basa en una arquitectura de red de nodos distribuidos que trabajan en conjunto para ofrecer un sistema de streaming de música robusto, escalable y altamente disponible. La propuesta consiste en la implementación de varios tipos de nodos, cada uno con roles y responsabilidades específicas, que trabajan de manera integrada para proporcionar un servicio de streaming eficiente y de alta calidad.

El sistema está compuesto inicialmente por los nodos Web, que son los encargados de ofrecer la interfaz gráfica a los usuarios. Estos nodos son la primera interacción del usuario con el sistema, proporcionando una interfaz fácil de usar y atractiva para subir, buscar, listar y reproducir música.

Para manejar las principales solicitudes y enrutarlas a las subredes apropiadas, tenemos los nodos API. Los nodos API actúan como un intermediario entre la interfaz de usuario y los demás nodos del sistema, procesando las solicitudes del usuario y enviándolas al nodo correcto para su ejecución.

A continuación, contamos con los nodos Peer y Tracker. Los nodos Peer tienen la responsabilidad de almacenar la música en el sistema. Almacenando los datos de las canciones de manera distribuida, Gotify garantiza un acceso rápido y eficiente a la música, así como una alta tolerancia a los fallos. Por otro lado, los nodos Tracker se encargan de almacenar el conjunto de metadatos asociados a cada pista de música. Esto permite un sistema de búsqueda eficiente y orientado, ya que los usuarios pueden buscar música basándose en metadatos como el género, el autor, el álbum y el nombre de la canción.

Finalmente, el sistema cuenta con nodos servidores DNS. Estos nodos juegan un papel esencial en la resolución de las solicitudes del navegador, proporcionando la instancia de un nodo Web que contiene la interfaz visual con la que un usuario interactuará con el sistema. Además, los nodos DNS resuelven las solicitudes desde un nodo Web hacia un nodo API que pueda responder a una solicitud específica.

En resumen, el proyecto Gotify se basa en una arquitectura de red de nodos distribuidos para proporcionar un sistema de streaming de música de alta calidad y fácil de usar, con una gran tolerancia a los fallos y una gran capacidad para manejar un gran volumen de usuarios y datos.

### Protocolo Kademlia:

Kademlia es un protocolo para redes peer-to-peer (P2P) que especifica la estructura de una red distribuida y la manera en que los nodos deben interactuar entre sí. Este protocolo se basa en una tabla hash distribuida (DHT), la cual permite que los nodos almacenen y recuperen información de manera eficiente en un entorno de red distribuida. Kademlia fue diseñado para ser altamente resistente a los fallos, con una capacidad superior para manejar entradas y salidas frecuentes de nodos en la red.

Kademlia se distingue por su algoritmo de enrutamiento que utiliza una métrica de distancia XOR. Este enfoque garantiza una alta eficiencia al minimizar el número de saltos requeridos para localizar un nodo en la red. El protocolo también implementa un sistema de 'k-buckets', que almacena información de contacto de otros nodos en la red, ayudando a mantener la red conectada y permitiendo un enrutamiento eficaz de las solicitudes.

Entre los principales métodos del protocolo Kademlia se encuentran: "STORE" para almacenar un par clave-valor en la red, "FIND_NODE" para encontrar la información de contacto de un nodo dado su ID, y "FIND_VALUE" para recuperar el valor asociado a una clave específica.

En el proyecto Gotify, utilizamos el protocolo Kademlia para el desarrollo de los nodos Tracker y Peer. Dada su eficiencia y resistencia a los fallos, Kademlia se presenta como una opción sólida para la implementación de nuestro sistema distribuido. Nos permite mantener una gran cantidad de datos distribuidos de manera eficiente y segura, lo cual es esencial para un sistema de streaming de música como Gotify.

Además, elegimos Kademlia sobre otros protocolos, como Chord, debido a varias razones. En primer lugar, Kademlia proporciona una mejor tolerancia a los fallos a través de su sistema de k-buckets, lo que permite al sistema recuperarse rápidamente de la pérdida de nodos. En segundo lugar, el algoritmo de enrutamiento basado en la métrica de distancia XOR de Kademlia proporciona una mejor eficiencia en comparación con el enfoque de anillos de Chord, lo que se traduce en una menor latencia y un rendimiento más rápido.

En resumen, la adopción del protocolo Kademlia proporciona al sistema Gotify un mecanismo robusto y eficiente para el almacenamiento y recuperación de datos, así como para el enrutamiento de solicitudes, lo que resulta esencial para mantener un servicio de streaming de música de alta calidad y con una gran tolerancia a los fallos.

### Web:
Los nodos Web en el sistema Gotify son la puerta de entrada para los usuarios humanos al sistema distribuido. Se encargan de proporcionar una interfaz gráfica intuitiva y fácil de usar que permite a los usuarios interactuar de manera transparente con el sistema distribuido subyacente. Estos nodos son responsables de traducir las acciones del usuario en solicitudes que el sistema distribuido pueda entender y procesar.

Un aspecto crucial de estos nodos Web es su papel en la tolerancia a fallos y en el equilibrio de la carga de la red. Para garantizar la alta disponibilidad del servicio y una experiencia de usuario fluida, se recomienda implementar varias instancias de nodos Web. De esta manera, si uno de ellos falla o está sobrecargado, los usuarios pueden ser redirigidos a otra instancia activa. Este enfoque distribuido ayuda a mantener la eficiencia y el rendimiento del sistema incluso ante un aumento de la demanda o fallos individuales.

Los nodos Web proporcionan varias funcionalidades clave para los usuarios. Estas incluyen la capacidad de subir música al sistema, reproducirla en streaming y aplicar filtros para listar las canciones disponibles. Esto significa que los usuarios pueden cargar su música preferida, reproducirla en cualquier momento y buscar fácilmente canciones en función de criterios específicos, como el género, el autor, el álbum o el nombre de la canción.

### Nodos DNS:
Los nodos DNS (Sistema de Nombres de Dominio) desempeñan un papel crítico en la infraestructura de Gotify, permitiendo la resolución de nombres de dominio en direcciones IP. La utilización de los nodos DNS es fundamental en dos momentos claves en el funcionamiento del sistema.

Primero, cuando un cliente desde su navegador intenta acceder al sistema Gotify a través de la URL "gotify.com". Aquí, los nodos DNS entran en acción, traduciendo esta URL en la dirección IP correspondiente al nodo Web activo en ese momento. Este proceso es esencial para la interacción inicial del usuario con el sistema, ya que asegura que el usuario pueda acceder de manera efectiva a la interfaz de Gotify a través de su navegador.

El segundo momento crucial en el que los nodos DNS intervienen es cuando el nodo Web, a través del cual el usuario está interactuando con el sistema, necesita comunicarse con un nodo API para procesar una solicitud específica del usuario. En este punto, los nodos DNS son responsables de traducir el nombre de dominio "api.gotify.com" en la dirección IP del nodo API adecuado.

Los nodos DNS logran estas tareas a través de un llamado de difusión o 'broadcast'. Este proceso permite a los nodos DNS descubrir los nodos Web y API activos en la red. De esta manera, pueden proporcionar direcciones IP precisas y actualizadas para las solicitudes entrantes, ya sea desde el navegador del usuario o desde el nodo Web.

### Nodos API:

Los nodos API en el sistema Gotify desempeñan un papel central en la coordinación de las operaciones fundamentales y en la facilitación de la comunicación entre las diversas partes del sistema.

Los nodos API proporcionan las siguientes funcionalidades clave:

1. **GetSongByID**: Esta función es responsable de recuperar la canción específica que se va a reproducir. Para ello, se comunica con los nodos Peer en los que se almacena la música. Una vez que se ha identificado la canción por su ID, el nodo API solicita la transferencia de datos de la canción desde el nodo Peer para su reproducción.

2. **FilterSongs**: Esta función se encarga de buscar canciones en función de criterios de filtrado específicos, como el género, el autor, el álbum, etc. Para realizar esta búsqueda, el nodo API consulta a los nodos Tracker, que mantienen el conjunto de metadatos asociados a cada canción.

3. **StoreSong**: Esta función permite almacenar una nueva canción en el sistema. Almacena la música en los nodos Peer para su reproducción posterior y los metadatos de la canción en los nodos Tracker para facilitar las búsquedas futuras.

Para facilitar estas funcionalidades, los nodos API están equipados con las capacidades necesarias para interactuar con los nodos Tracker y Peer como si fueran parte de su propia red. En una posible implementación, un nodo API podría incluir un nodo Peer y un nodo Tracker "ligeros" que, si bien no almacenan información, facilitan la comunicación con otros nodos Tracker y Peer en la red. Esto permitiría al nodo API realizar consultas a nivel de red a estos nodos de manera eficiente.

En resumen, los nodos API en Gotify son esenciales para el funcionamiento del sistema, ya que proporcionan una capa de abstracción entre la interfaz del usuario y la complejidad subyacente del sistema distribuido. Al proporcionar una interfaz de comunicación estandarizada, los nodos API permiten la integración eficiente de las diferentes partes del sistema y facilitan la implementación de funcionalidades clave para el streaming de música.

### Nodos Peer:

Los nodos Peer en Gotify son una parte integral de la red Kademlia y juegan un papel vital en el almacenamiento y la recuperación de la música. Estos nodos utilizan los métodos fundamentales de Kademlia y proporcionan funcionalidades clave para el streaming de música.

Las funcionalidades clave proporcionadas por los nodos Peer son:

1. **StoreSong**: Esta función es responsable de almacenar la música en los nodos Peer. Para hacer esto de manera eficiente, cada nodo Peer tiene una base de datos asociada. Esta base de datos almacena los datos de las canciones de manera eficiente para permitir un almacenamiento robusto y una recuperación rápida de la música. Cada canción almacenada en un nodo Peer se identifica de manera única para permitir la recuperación precisa.

2. **GetSong**: Esta función permite recuperar fragmentos o 'chunks' de bytes de una canción específica almacenada en el nodo Peer. El objetivo de esta funcionalidad es permitir el streaming de estos fragmentos de canción en el componente de reproducción de la interfaz web. Este método permite la reproducción de música en tiempo real, esencial para la experiencia de usuario en Gotify.

En resumen, los nodos Peer son fundamentales para la capacidad de Gotify de almacenar y transmitir música. Al aprovechar la arquitectura de la red Kademlia, los nodos Peer proporcionan un sistema eficiente y escalable para el almacenamiento de música, asegurando que Gotify pueda manejar un amplio rango de música y usuarios sin comprometer la velocidad o la calidad del streaming.

### Nodos Tracker:
Los nodos Tracker en Gotify forman una red basada en el protocolo Kademlia, desempeñando un papel crítico en la gestión de metadatos de canciones y la facilitación de funciones de búsqueda y listado de canciones. Su diseño y funcionalidad están optimizados para responder a consultas basadas en metadatos, lo que permite a los usuarios buscar y descubrir música de manera eficiente en el sistema.

Los nodos Tracker proporcionan las siguientes funcionalidades clave:

1. **StoreSongMetadata**: Esta función es responsable de almacenar los metadatos de una canción en la red de nodos Tracker. Los metadatos incluyen información sobre el título, artista, género y álbum de la canción. Además de almacenar estos metadatos, la función también almacena todas las posibles combinaciones de consultas que podrían surgir a partir de estos metadatos. Este enfoque permite que la red de nodos Tracker responda de manera eficiente a diversas consultas de filtrado y listado de canciones.

2. **Consulta y Recuperación de Canciones**: En su almacenamiento interno, cada nodo Tracker mantiene un índice de las consultas de metadatos y las canciones que coinciden con estas consultas. Para cada consulta, se guarda una lista de las canciones que satisfacen la consulta. Además, cada canción se almacena con su conjunto completo de metadatos y un identificador único. Este identificador es crítico ya que permite a los nodos Tracker referenciar la canción correspondiente en la red de nodos Peer.

En resumen, los nodos Tracker son vitales para la experiencia del usuario en Gotify, facilitando la búsqueda y descubrimiento de música a través de una gestión eficiente de los metadatos. Aprovechando la arquitectura de la red Kademlia, los nodos Tracker proporcionan una solución escalable y robusta para el almacenamiento y consulta de metadatos de canciones.

### Flujo de Funcionamiento del Proyecto:
En esta sección, vamos a describir el flujo de las principales operaciones del sistema: guardar una canción, filtrar las canciones y reproducirlas.

1. **Guardando una Canción**: 

   El flujo para guardar una canción en el sistema Gotify comienza en la interfaz web. Un usuario interactúa con esta interfaz para subir una canción, proporcionando el archivo de música y los metadatos relevantes (título, artista, género y álbum). 

   Una vez que el usuario ha proporcionado toda la información necesaria, la interfaz web envía una solicitud a los nodos API utilizando la función StoreSong. Aquí es donde interviene el nodo DNS, resolviendo la dirección del nodo API apropiado para esta solicitud.

   Los nodos API, utilizando su funcionalidad interna de red Kademlia, se comunican con los nodos Peer y Tracker para almacenar la canción y sus metadatos, respectivamente. Los nodos Peer almacenan la canción, mientras que los nodos Tracker almacenan los metadatos y todas las posibles combinaciones de consultas relacionadas.

2. **Filtrado de Canciones**: 

   El filtrado de canciones también comienza en la interfaz web. Los usuarios pueden buscar canciones por título, artista, género y álbum. Esta consulta se envía a los nodos API a través de la función FilterSongs.

   Los nodos API, con la ayuda del nodo DNS para resolver las direcciones necesarias, consultan a los nodos Tracker con la consulta de filtrado. Los nodos Tracker responden con una lista de canciones que cumplen con los criterios de búsqueda, que luego se devuelve a la interfaz web para que el usuario pueda ver los resultados.

3. **Reproducción de Canciones**:

   La reproducción de canciones comienza cuando un usuario selecciona una canción de la lista de resultados del filtrado. La interfaz web envía una solicitud a los nodos API con la identificación de la canción utilizando la función GetSongByID.

   Los nodos API, nuevamente con la ayuda del nodo DNS, consultan a los nodos Peer con la identificación de la canción. Los nodos Peer recuperan la canción y envían los datos al nodo API, que luego se transmiten a la interfaz web para su reproducción. 

Cada una de estas operaciones depende de la colaboración efectiva de todas las partes del sistema. Juntos, los nodos DNS, Web, API, Peer y Tracker permiten una experiencia de usuario fluida y eficiente, desde el almacenamiento hasta la reproducción de música.

### Tolerancia a Fallas:
En un entorno distribuido como el que presenta Gotify, la tolerancia a fallas es una consideración de diseño crítica. Este sistema fue diseñado para manejar y recuperarse de una variedad de fallas, garantizando la disponibilidad y fiabilidad del servicio en todo momento.

1. **Disponibilidad de Nodos**: El diseño del sistema asegura que, siempre que exista al menos un nodo de cada tipo (Web, DNS, API, Peer, y Tracker), la red sea capaz de resolver las distintas peticiones de los usuarios. Esta característica proporciona una resistencia fundamental contra las interrupciones del servicio, ya que incluso la pérdida de nodos individuales no impide la capacidad de la red para funcionar.

2. **Replicación de Datos**: Cada nodo de la red, ya sea un Peer o un Tracker, se encarga de replicar la información a intervalos regulares mediante el método 'Republish'. Esta replicación garantiza que, aunque los nodos que originalmente contenían una pieza específica de información puedan caer, la información misma se propaga y se mantiene en la red con el tiempo. Esta característica aumenta significativamente la resiliencia de la red contra la pérdida de datos.

3. **Unión a la Red**: Los nodos Peer y Tracker implementan una función 'JoinNetwork' al unirse a la red. Esta función busca nodos bootstrap de su subred para establecer una conexión inicial. Los nodos bootstrap son específicos para cada subred de Peers y Trackers, y actúan como nodos de entrada para nuevos nodos que se unen a la red. Además, en el caso de una partición de la red (la red pierde conexidad), los nodos bootstrap facilitan la reconexión y la continuidad de la red.

4. **Gestión de la Conexión de Reproducción**: En el caso extremo de que se pierda la conexión con un nodo Peer que está reproduciendo una canción, el nodo API que solicitó la canción mantiene un registro de la sección o chunk de la canción que debe buscar a continuación para la reproducción. Este sistema asegura la continuidad en la reproducción de música, incluso si se pierde la conexión con un nodo durante la reproducción.

La tolerancia a fallas en Gotify se logra mediante la combinación de múltiples estrategias y técnicas, que trabajan juntas para manejar tanto las fallas esperadas como las inesperadas. A través de la replicación de datos, la presencia de nodos bootstrap y la gestión efectiva de la conexión de reproducción, Gotify proporciona un servicio robusto y confiable que puede adaptarse y recuperarse de una variedad de escenarios de fallas.

## Conclusiones

El proyecto Gotify representa un innovador enfoque para la creación de un sistema de streaming de música distribuido. A través de la utilización de una variedad de nodos con roles especializados, incluyendo nodos Web, DNS, API, Peer, y Tracker, Gotify logra ofrecer un sistema capaz de proporcionar servicios de streaming eficientes y confiables.

En su diseño, Gotify destaca por la implementación del protocolo Kademlia, logrando una estructura de red descentralizada que promueve la resistencia y la eficiencia en la búsqueda y almacenamiento de datos. Además, las características inherentes de Kademlia proporcionan a Gotify una sólida tolerancia a fallas, asegurando que el sistema pueda adaptarse y recuperarse de diversas situaciones adversas.

Cada componente del sistema Gotify desempeña un papel esencial en la provisión del servicio, desde la interfaz de usuario hasta la recuperación y almacenamiento de datos. El uso de nodos API para enlazar la red con la interfaz del usuario, y el empleo de nodos Peer y Tracker para el almacenamiento y recuperación de música y metadatos, son ejemplos de cómo Gotify aprovecha los principios de los sistemas distribuidos para ofrecer un servicio de alta calidad.

En última instancia, Gotify demuestra cómo los principios de los sistemas distribuidos se pueden aplicar de manera efectiva para desarrollar una plataforma de streaming de música que es resistente, escalable y capaz de proporcionar una experiencia de usuario superior. Este proyecto sirve como un valioso caso de estudio de cómo los avances en la tecnología de sistemas distribuidos pueden ser utilizados para mejorar y optimizar la entrega de servicios digitales en la vida cotidiana.