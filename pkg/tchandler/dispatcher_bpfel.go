// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || arm || arm64 || loong64 || mips64le || mipsle || ppc64le || riscv64

package tchandler

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

// loadDispatcher returns the embedded CollectionSpec for dispatcher.
func loadDispatcher() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_DispatcherBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load dispatcher: %w", err)
	}

	return spec, err
}

// loadDispatcherObjects loads dispatcher and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*dispatcherObjects
//	*dispatcherPrograms
//	*dispatcherMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadDispatcherObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadDispatcher()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// dispatcherSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type dispatcherSpecs struct {
	dispatcherProgramSpecs
	dispatcherMapSpecs
	dispatcherVariableSpecs
}

// dispatcherProgramSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type dispatcherProgramSpecs struct {
	IgNetDisp *ebpf.ProgramSpec `ebpf:"ig_net_disp"`
}

// dispatcherMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type dispatcherMapSpecs struct {
	GadgetTailCall *ebpf.MapSpec `ebpf:"gadget_tail_call"`
}

// dispatcherVariableSpecs contains global variables before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type dispatcherVariableSpecs struct {
	CurrentNetns *ebpf.VariableSpec `ebpf:"current_netns"`
}

// dispatcherObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadDispatcherObjects or ebpf.CollectionSpec.LoadAndAssign.
type dispatcherObjects struct {
	dispatcherPrograms
	dispatcherMaps
	dispatcherVariables
}

func (o *dispatcherObjects) Close() error {
	return _DispatcherClose(
		&o.dispatcherPrograms,
		&o.dispatcherMaps,
	)
}

// dispatcherMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadDispatcherObjects or ebpf.CollectionSpec.LoadAndAssign.
type dispatcherMaps struct {
	GadgetTailCall *ebpf.Map `ebpf:"gadget_tail_call"`
}

func (m *dispatcherMaps) Close() error {
	return _DispatcherClose(
		m.GadgetTailCall,
	)
}

// dispatcherVariables contains all global variables after they have been loaded into the kernel.
//
// It can be passed to loadDispatcherObjects or ebpf.CollectionSpec.LoadAndAssign.
type dispatcherVariables struct {
	CurrentNetns *ebpf.Variable `ebpf:"current_netns"`
}

// dispatcherPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadDispatcherObjects or ebpf.CollectionSpec.LoadAndAssign.
type dispatcherPrograms struct {
	IgNetDisp *ebpf.Program `ebpf:"ig_net_disp"`
}

func (p *dispatcherPrograms) Close() error {
	return _DispatcherClose(
		p.IgNetDisp,
	)
}

func _DispatcherClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed dispatcher_bpfel.o
var _DispatcherBytes []byte
