
# Sessionbeat

Welcome to Sessionbeat. / Bienvenidos al Sessionbeat


## Overview  / Vista general

![EN](https://cdn1.iconfinder.com/data/icons/famfamfam_flag_icons/gb.png) : SessionBeat is a beat that allows you to monitor remote desktop sessions of a Windows Server using the RDP protocol, this information is consulted using the `PowerShell module PSTerminalServices`, we plan that future versions make the query directly to the `WMI`.

![ES](https://cdn1.iconfinder.com/data/icons/famfamfam_flag_icons/es.png) : SessionBeat es un beat  que te permite monitorear las sessiones de escritorios remotos de un servidor Windows mediante el protocolo RDP, esta informacion es consultada usando el modulo de `PowerShell PSTerminalServices`, planeamos que las versiones futuras realicen la consulta de manera directa a la `WMI`.

## Requirements / Requerimientos

* [Golang](https://golang.org/dl/) 1.7 (Only if you want to modify the code)
* [PSTerminalServices](https://archive.codeplex.com/?p=psterminalservices)

As we have mentioned before, the beat depends of `PowerShell PSTerminalServices`, thus you need import the module, you can follow thes [tutorial](https://archive.codeplex.com/?p=psterminalservices).

Como mencionamos anteriormente, el beat depende  de `PowerShell PSTerminalServices`, por lo que necesita importar el módulo, puede seguir este [tutorial](https://archive.codeplex.com/?p=psterminalservices).

**Remember import and try the module / Recuerda importar y probar el módulo**

    Import-Module PSTerminalServices
    Get-TSSession


### Fields / Campos

* SessionId: Each session is identified by a unique session ID, assigned when a user logs on to a Remote Desktop Services / Cada sesión se identifica mediante un ID de sesión único, que se asigna cuando un usuario inicia sesión en los Servicios de escritorio remoto
* IPAddress: IP Address of user / Dirección IP del usuario
* ClientName: The name of the device from where do the connection / El nombre del dispositivo desde donde se realiza la conexión
* State: The state of the session, could be (Active, Listening, Disconnected, Connected, ConnectQuery, Down) // El estado de la sesión, puede ser (Active, Listening, Disconnected, Connected, ConnectQuery, Down)
* Username: User name to which the session is connected / Nombre de usuario al que está conectado la session



# Sessionbeat

## Getting Started with Sessionbeat

### Requirements

* [Golang](https://golang.org/dl/) 1.7

### Init Project
To get running with Sessionbeat and also install the
dependencies, run the following command:

```
make setup
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

To push Sessionbeat in the git repository, run the following commands:

```
git remote set-url origin https://github.com/3AT-Development/SessionsBeat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for Sessionbeat run the command below. This will generate a binary
in the same directory with the name sessionbeat.

```
make
```


### Run

To run Sessionbeat with debugging output enabled, run:

```
./sessionbeat -c sessionbeat.yml -e -d "*"
```


### Test

To test Sessionbeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `fields.yml` by running the following command.

```
make update
```


### Cleanup

To clean  Sessionbeat source code, run the following commands:

```
make fmt
make simplify
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone Sessionbeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/src/github.com/3at-developtment/sessionbeat
git clone https://github.com/3AT-Development/SessionsBeat ${GOPATH}/src/github.com/3at-developtment/sessionbeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).
